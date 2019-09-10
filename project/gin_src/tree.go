// Copyright 2013 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// at https://github.com/julienschmidt/httprouter/blob/master/LICENSE

package gin

import (
	"net/url"
	"strings"
	"unicode"
)

// Param is a single URL parameter, consisting of a key and a value.
type Param struct {
	Key   string
	Value string
}

// Params is a Param-slice, as returned by the router.
// The slice is ordered, the first URL parameter is also the first slice value.
// It is therefore safe to read values by the index.
type Params []Param

// Get returns the value of the first Param which key matches the given name.
// If no matching Param is found, an empty string is returned.
func (ps Params) Get(name string) (string, bool) {
	for _, entry := range ps {
		if entry.Key == name {
			return entry.Value, true
		}
	}
	return "", false
}

// ByName returns the value of the first Param which key matches the given name.
// If no matching Param is found, an empty string is returned.
func (ps Params) ByName(name string) (va string) {
	va, _ = ps.Get(name)
	return
}

type methodTree struct {
	method string
	root   *node
}

type methodTrees []methodTree

func (trees methodTrees) get(method string) *node {
	for _, tree := range trees {
		if tree.method == method {
			return tree.root
		}
	}
	return nil
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// 计算路径中有多少参数，其实就是冒号与星号的个数。
func countParams(path string) uint8 {
	var n uint
	for i := 0; i < len(path); i++ {
		if path[i] != ':' && path[i] != '*' {
			continue
		}
		n++
	}
	if n >= 255 {
		return 255
	}
	return uint8(n)
}

type nodeType uint8

const (
	static nodeType = iota // default
	root
	param
	catchAll
)

type node struct {
	path      string 			// 当前节点对应的路径中的字符串
	indices   string 			// 子节点索引，当子节点为非参数类型，即本节点的wildChild为false时，会将每个子节点的首字母放在该索引数组。说是数组，实际上是个string。
	children  []*node   		// 当前节点的所有直接子节点
	handlers  HandlersChain  	// 该路径对应要执行的的函数
	priority  uint32 			// 优先级，查找的时候会用到,表示当前节点加上所有子节点的数目
	nType     nodeType 			// 当前节点类型，有四个枚举值: 分别为 static/root/param/catchAll。
								// static 非根节点的普通字符串节点, root 根节点, param 参数节点，例如 :id, catchAll 通配符节点，例如 *anyway
	maxParams uint8 			// path 中的参数最大数量，最大只能保存 255 个（超过这个的情况貌似太难见到了）
   			 					// 这里是一个非负的 8 进制数字，最大也只能是 255 了
	wildChild bool  			// 判断当前节点路径是不是含有参数的节点
}

// increments priority of the given child and reorders if necessary.
func (n *node) incrementChildPrio(pos int) int {
	n.children[pos].priority++
	prio := n.children[pos].priority

	// adjust position (move to front)
	newPos := pos
	for newPos > 0 && n.children[newPos-1].priority < prio {
		// swap node positions
		n.children[newPos-1], n.children[newPos] = n.children[newPos], n.children[newPos-1]

		newPos--
	}

	// build new index char string
	if newPos != pos {
		n.indices = n.indices[:newPos] + // unchanged prefix, might be empty
			n.indices[pos:pos+1] + // the index char we move
			n.indices[newPos:pos] + n.indices[pos+1:] // rest without char at 'pos'
	}

	return newPos
}

// addRoute adds a node with the given handle to the path.
// Not concurrency-safe!
// addRoute添加节点
func (n *node) addRoute(path string, handlers HandlersChain) {
	fullPath := path
	// 请求到这个方法，就给当前节点的权重 + 1
	n.priority++
	// 计算传入的路径参数的数量
	numParams := countParams(path)

	// 如果树不是空的
	// 判断的条件是当前节点的 path 的字符串长度和子节点的数量全部都大于 0
	// 就是说如果当前节点是一个空的节点，或者当前的节点是一个叶子节点，就直接
    // 进入 else 在当前节点下面添加子节点
	if len(n.path) > 0 || len(n.children) > 0 {
	// 定义一个 lable ，循环里面可以直接 break 到这里，适合这种嵌套的比较深的
	walk:
		for {
			// 如果传入的节点的最大参数的数量大于当前节点记录的数量，替换
			// 更新当前node的最大参数个数
			if numParams > n.maxParams {
				n.maxParams = numParams
			}

			// 找到最长公共前缀
			// 公共前缀不包含 ":" 或 "*"
			i := 0
			// 将最大值设置成长度较小的路径的长度
			max := min(len(path), len(n.path))
			// 循环，计算出当前节点和添加的节点共同前缀的长度
			for i < max && path[i] == n.path[i] {
				i++
			}

			// 如果相同前缀的长度比当前节点保存的 path 短
            // 比如当前节点现在的 path 是 sup ，添加的节点的 path 是 search
            // 它们相同的前缀就变成了 s ， s 比 sup 要短，符合 if 的条件，要做处理
			if i < len(n.path) {
				// 将当前节点的属性定义到一个子节点中，没有注释的属性不变，保持原样
				child := node{
					path:      n.path[i:],  		// path 是当前节点的 path 去除公共前缀长度的部分
					wildChild: n.wildChild,
					indices:   n.indices,
					children:  n.children,
					handlers:  n.handlers,
					priority:  n.priority - 1, 		 // 权重 -1 
				}

				// 遍历当前节点的所有子节点（当前节点变成子节点之后的节点），
                // 如果最大参数数量大于当前节点的数量，更新
				for i := range child.children {
					if child.children[i].maxParams > child.maxParams {
						child.maxParams = child.children[i].maxParams
					}
				}

				// 在当前节点的子节点定义为当前节点转换后的子节点
				n.children = []*node{&child}
				// 获取子节点的首字母,因为上面分割的时候是从 i 的位置开始分割
                // 所以 n.path[i] 可以去除子节点的首字母，理论上去 child.path[0] 也是可以的
                // 这里的 n.path[i] 取出来的是一个 uint8 类型的数字（代表字符），
                // 先用 []byte 包装一下数字再转换成字符串格式
				n.indices = string([]byte{n.path[i]})
				// 更新当前节点的 path 为新的公共前缀
				n.path = path[:i]
				// 将 handle 设置为 nil
				n.handlers = nil
				// 肯定没有参数了，已经变成了一个没有 handle 的节点了
				n.wildChild = false
			}

			// 将新的节点添加到此节点的子节点， 这里是新添加节点的子节点
			if i < len(path) {
				// 截取掉公共部分，剩余的是子节点
				path = path[i:]

				// 如果当前路径有参数
                // 如果进入了上面 if i < len(n.path) 这个条件，这里就不会成立了
                // 因为上一个 if 中将 n.wildChild 重新定义成了 false 
                // 什么情况会进入到这里呢 ? 
                // 1. 上面的 if 不生效，也就是说不会有新的公共前缀， n.path = i 的时候
                // 2. 当前节点的 path 是一个参数节点就是像这种的 :post
                // 就是定义路由时候是这种形式的： blog/:post/update
				if n.wildChild {
					n = n.children[0]
					n.priority++

					// Update maxParams of the child node
					if numParams > n.maxParams {
						n.maxParams = numParams
					}
					numParams--

					// Check if the wildcard matches
					if len(path) >= len(n.path) && n.path == path[:len(n.path)] {
						// check for longer wildcard, e.g. :name and :names
						if len(n.path) >= len(path) || path[len(n.path)] == '/' {
							continue walk
						}
					}

					pathSeg := path
					if n.nType != catchAll {
						pathSeg = strings.SplitN(path, "/", 2)[0]
					}
					prefix := fullPath[:strings.Index(fullPath, pathSeg)] + n.path
					panic("'" + pathSeg +
						"' in new path '" + fullPath +
						"' conflicts with existing wildcard '" + n.path +
						"' in existing prefix '" + prefix +
						"'")
				}

				c := path[0]

				// slash after param
				if n.nType == param && c == '/' && len(n.children) == 1 {
					n = n.children[0]
					n.priority++
					continue walk
				}

				// Check if a child with the next path byte exists
				for i := 0; i < len(n.indices); i++ {
					if c == n.indices[i] {
						i = n.incrementChildPrio(i)
						n = n.children[i]
						continue walk
					}
				}

				// Otherwise insert it
				if c != ':' && c != '*' {
					// []byte for proper unicode char conversion, see #65
					n.indices += string([]byte{c})
					child := &node{
						maxParams: numParams,
					}
					n.children = append(n.children, child)
					n.incrementChildPrio(len(n.indices) - 1)
					n = child
				}
				n.insertChild(numParams, path, fullPath, handlers)
				return

			} else if i == len(path) { // Make node a (in-path) leaf
				if n.handlers != nil {
					panic("handlers are already registered for path '" + fullPath + "'")
				}
				n.handlers = handlers
			}
			return
		}
	} else { // Empty tree
		//如果节点为空，将插入的节点类型定义为root节点，就是根节点
		n.insertChild(numParams, path, fullPath, handlers)
		n.nType = root
	}
}


// numParams 参数个数
// path 插入的子节点的路径
// fullPath 完整路径，就是注册路由时候的路径，没有被处理过的
// 注册路由对应的 handle 函数
func (n *node) insertChild(numParams uint8, path string, fullPath string, handlers HandlersChain) {
	var offset int // 已经处理过的路径的所有字节数

	// 查找前缀，知道第一个通配符（ 以 ':' 或 '*' 开头
    // 就是要将 path 遍历，提取出参数
    // 只要不是通配符开头的就不做处理，证明这个路由是没有参数的路由
	for i, max := 0, len(path); numParams > 0; i++ {
		c := path[i]
		// 如果不是 : 或 * 跳过本次循环，不做任何处理
		if c != ':' && c != '*' {
			continue
		}

		// 查询通配符后面的字符，直到查到 '/' 或者结束
		end := i + 1
		for end < max && path[end] != '/' {
			switch path[end] {
			// 通配符后面的名称不能包含 : 或 * ， 如 ::name 或 :*name 不允许定义
			case ':', '*':
				panic("only one wildcard per path segment is allowed, has: '" +
					path[i:] + "' in path '" + fullPath + "'")
			default:
				end++
			}
		}

		// 检查通配符所在的位置，是否已经有子节点，如果有，就不能再插入
        // 例如： 已经定义了 /hello/name ， 就不能再定义 /hello/:param
		if len(n.children) > 0 {
			panic("wildcard route '" + path[i:end] +
				"' conflicts with existing children in path '" + fullPath + "'")
		}

		// 检查通配符是否有一个名字
        // 上面定义 end = i+1 ， 后面的 for 又执行了 ++ 操作，所以通配符 : 或 * 后面最少
        // 要有一个字符, 如： :a 或 :name ， :a 的时候 end 就是 i+2
        // 所以如果 end - i < 2 ，就是通配符后面没有对应的名称， 就会 panic
		if end-i < 2 {
			panic("wildcards must be named with a non-empty name in path '" + fullPath + "'")
		}

		// 如果 c 是 : 通配符的时候
		if c == ':' { // param
			// 从 offset 的位置，到查询到通配符的位置分割 path
            // 并把分割出来的路径定义到节点的 path 属性
			if i > 0 {
				// 跟路径不包含:
				n.path = path[offset:i]
				// 开始的位置变成了通配符所在的位置
				offset = i
			}

			// 将参数部分定义成一个子节点
			child := &node{
				nType:     param,
				maxParams: numParams,
			}
			// 用新定义的子节点初始化一个 children 属性
			n.children = []*node{child}
			n.wildChild = true      		// 标记上当前这个节点是一个包含参数的节点的节点
			// 将新创建的节点定义为当前节点，这个要想一下，到这里这种操作已经有不少了
            // 因为一直都是指针操作，修改都是指针的引用，所以定义好的层级关系不会被改变
			n = child 			
			// 新的节点权重 +1			
			n.priority++
			// 最大参数个数 - 1
			numParams--         

			// if the path doesn't end with the wildcard, then there
			// will be another non-wildcard subpath starting with '/'
			// 这个 end 有可能是结束或者下一个 / 的位置
            // 如果小于路径的最大长度，代表还包含子路径（也就是说后面还有子节点）
			if end < max {
				//设置参数节点的path
				n.path = path[offset:end]
				offset = end

				child := &node{
					maxParams: numParams,
					priority:  1,
				}
				n.children = []*node{child}
				n = child
			}

		} else { // catchAll
			// 这里的意思是， * 匹配的路径只允许定义在路由的最后一部分 
            // 比如 : /hello/*world 是允许的， /hello/*world/more 这种就会 painc
            // 这种路径就是会将 hello/ 后面的所有内容变成 world 的变量
            // 比如地址栏输入： /hello/one/two/more ，获取到的参数 world = one/twq/more
            // 不会再将后面的 / 作为路径处理了
			if end != max || numParams > 1 {
				panic("catch-all routes are only allowed at the end of the path in path '" + fullPath + "'")
			}

			//路径结尾不能为/
			if len(n.path) > 0 && n.path[len(n.path)-1] == '/' {
				panic("catch-all conflicts with existing handle for the path segment root in path '" + fullPath + "'")
			}

			// catchAll路由前面的字符必须为/
			i--
			if path[i] != '/' {
				panic("no / before catch-all in path '" + fullPath + "'")
			}

			n.path = path[offset:i]

			// first node: catchAll node with empty path
			child := &node{
				wildChild: true,
				nType:     catchAll,
				maxParams: 1,
			}
			n.children = []*node{child}
			n.indices = string(path[i])
			n = child
			n.priority++

			// second node: node holding the variable
			child = &node{
				path:      path[i:],
				nType:     catchAll,
				maxParams: 1,
				handlers:  handlers,
				priority:  1,
			}
			n.children = []*node{child}

			return
		}
	}

	// insert remaining path part and handle to the leaf
	n.path = path[offset:]
	n.handlers = handlers
}

// getValue returns the handle registered with the given path (key). The values of
// wildcards are saved to a map.
// If no handle can be found, a TSR (trailing slash redirect) recommendation is
// made if a handle exists with an extra (without the) trailing slash for the
// given path.
func (n *node) getValue(path string, po Params, unescape bool) (handlers HandlersChain, p Params, tsr bool) {
	p = po
walk: // Outer loop for walking the tree
	for {
		if len(path) > len(n.path) {
			if path[:len(n.path)] == n.path {
				path = path[len(n.path):]
				// If this node does not have a wildcard (param or catchAll)
				// child,  we can just look up the next child node and continue
				// to walk down the tree
				if !n.wildChild {
					c := path[0]
					for i := 0; i < len(n.indices); i++ {
						if c == n.indices[i] {
							n = n.children[i]
							continue walk
						}
					}

					// Nothing found.
					// We can recommend to redirect to the same URL without a
					// trailing slash if a leaf exists for that path.
					tsr = path == "/" && n.handlers != nil
					return
				}

				// handle wildcard child
				n = n.children[0]
				switch n.nType {
				case param:
					// find param end (either '/' or path end)
					end := 0
					for end < len(path) && path[end] != '/' {
						end++
					}

					// save param value
					if cap(p) < int(n.maxParams) {
						p = make(Params, 0, n.maxParams)
					}
					i := len(p)
					p = p[:i+1] // expand slice within preallocated capacity
					p[i].Key = n.path[1:]
					val := path[:end]
					if unescape {
						var err error
						if p[i].Value, err = url.QueryUnescape(val); err != nil {
							p[i].Value = val // fallback, in case of error
						}
					} else {
						p[i].Value = val
					}

					// we need to go deeper!
					if end < len(path) {
						if len(n.children) > 0 {
							path = path[end:]
							n = n.children[0]
							continue walk
						}

						// ... but we can't
						tsr = len(path) == end+1
						return
					}

					if handlers = n.handlers; handlers != nil {
						return
					}
					if len(n.children) == 1 {
						// No handle found. Check if a handle for this path + a
						// trailing slash exists for TSR recommendation
						n = n.children[0]
						tsr = n.path == "/" && n.handlers != nil
					}

					return

				case catchAll:
					// save param value
					if cap(p) < int(n.maxParams) {
						p = make(Params, 0, n.maxParams)
					}
					i := len(p)
					p = p[:i+1] // expand slice within preallocated capacity
					p[i].Key = n.path[2:]
					if unescape {
						var err error
						if p[i].Value, err = url.QueryUnescape(path); err != nil {
							p[i].Value = path // fallback, in case of error
						}
					} else {
						p[i].Value = path
					}

					handlers = n.handlers
					return

				default:
					panic("invalid node type")
				}
			}
		} else if path == n.path {
			//如果当前path与当前结点的path相同

			//如果当前节点的handlers，直接返回
			if handlers = n.handlers; handlers != nil {
				return
			}

			//如果当前path=='/' 并且当前节点有参数子节点，并且当前节点不是根节点 /:id
			if path == "/" && n.wildChild && n.nType != root {
				tsr = true
				return
			}

			// No handle found. Check if a handle for this path + a
			// trailing slash exists for trailing slash recommendation
			for i := 0; i < len(n.indices); i++ {
				if n.indices[i] == '/' {
					n = n.children[i]
					tsr = (len(n.path) == 1 && n.handlers != nil) ||
						(n.nType == catchAll && n.children[0].handlers != nil)
					return
				}
			}

			return
		}

		// Nothing found. We can recommend to redirect to the same URL with an
		// extra trailing slash if a leaf exists for that path
		tsr = (path == "/") ||
			(len(n.path) == len(path)+1 && n.path[len(path)] == '/' &&
				path == n.path[:len(n.path)-1] && n.handlers != nil)
		return
	}
}

// findCaseInsensitivePath makes a case-insensitive lookup of the given path and tries to find a handler.
// It can optionally also fix trailing slashes.
// It returns the case-corrected path and a bool indicating whether the lookup
// was successful.
func (n *node) findCaseInsensitivePath(path string, fixTrailingSlash bool) (ciPath []byte, found bool) {
	ciPath = make([]byte, 0, len(path)+1) // preallocate enough memory

	// Outer loop for walking the tree
	for len(path) >= len(n.path) && strings.ToLower(path[:len(n.path)]) == strings.ToLower(n.path) {
		path = path[len(n.path):]
		ciPath = append(ciPath, n.path...)

		if len(path) > 0 {
			// If this node does not have a wildcard (param or catchAll) child,
			// we can just look up the next child node and continue to walk down
			// the tree
			if !n.wildChild {
				r := unicode.ToLower(rune(path[0]))
				for i, index := range n.indices {
					// must use recursive approach since both index and
					// ToLower(index) could exist. We must check both.
					if r == unicode.ToLower(index) {
						out, found := n.children[i].findCaseInsensitivePath(path, fixTrailingSlash)
						if found {
							return append(ciPath, out...), true
						}
					}
				}

				// Nothing found. We can recommend to redirect to the same URL
				// without a trailing slash if a leaf exists for that path
				found = fixTrailingSlash && path == "/" && n.handlers != nil
				return
			}

			n = n.children[0]
			switch n.nType {
			case param:
				// find param end (either '/' or path end)
				k := 0
				for k < len(path) && path[k] != '/' {
					k++
				}

				// add param value to case insensitive path
				ciPath = append(ciPath, path[:k]...)

				// we need to go deeper!
				if k < len(path) {
					if len(n.children) > 0 {
						path = path[k:]
						n = n.children[0]
						continue
					}

					// ... but we can't
					if fixTrailingSlash && len(path) == k+1 {
						return ciPath, true
					}
					return
				}

				if n.handlers != nil {
					return ciPath, true
				} else if fixTrailingSlash && len(n.children) == 1 {
					// No handle found. Check if a handle for this path + a
					// trailing slash exists
					n = n.children[0]
					if n.path == "/" && n.handlers != nil {
						return append(ciPath, '/'), true
					}
				}
				return

			case catchAll:
				return append(ciPath, path...), true

			default:
				panic("invalid node type")
			}
		} else {
			// We should have reached the node containing the handle.
			// Check if this node has a handle registered.
			if n.handlers != nil {
				return ciPath, true
			}

			// No handle found.
			// Try to fix the path by adding a trailing slash
			if fixTrailingSlash {
				for i := 0; i < len(n.indices); i++ {
					if n.indices[i] == '/' {
						n = n.children[i]
						if (len(n.path) == 1 && n.handlers != nil) ||
							(n.nType == catchAll && n.children[0].handlers != nil) {
							return append(ciPath, '/'), true
						}
						return
					}
				}
			}
			return
		}
	}

	// Nothing found.
	// Try to fix the path by adding / removing a trailing slash
	if fixTrailingSlash {
		if path == "/" {
			return ciPath, true
		}
		if len(path)+1 == len(n.path) && n.path[len(path)] == '/' &&
			strings.ToLower(path) == strings.ToLower(n.path[:len(path)]) &&
			n.handlers != nil {
			return append(ciPath, n.path...), true
		}
	}
	return
}
