# Golang Practice on Data Structure

这是一个使用 Go 语言实现的数据结构练习项目，包含多种常见的数据结构泛型实现及部分测试代码。`interface{}` 实现请查看 [go1.17](https://github.com/dreamy-xay/golang-data-structure/tree/go1.17) 分支。

## 目录结构

- **avltree**：AVL 树相关实现及测试代码。

  - `avltree.go`：AVL 树基础实现。
  - `avltree_map.go`：AVL 树映射实现。
  - `avltree_test.go`：AVL 树测试代码。
  - `multi_avltree.go`：多 AVL 树实现。
  - `multi_avltree_map.go`：多 AVL 树映射实现。

- **common**：通用功能和接口定义。

  - `function.go`：通用函数实现。
  - `interface.go`：通用接口定义。
  - `type.go`：通用类型定义。

- **hashmap**：哈希表相关实现及测试代码。

  - `hashmap.go`：哈希表基础实现。
  - `hashmap_test.go`：哈希表测试代码。

- **hashset**：哈希集合相关实现。

  - `hashset.go`：哈希集合基础实现。

- **list**：链表相关实现。

  - `list.go`：链表基础实现。

- **priorityqueue**：优先队列相关实现。

  - `priorityqueue.go`：优先队列基础实现。

- **queue**：队列相关实现。

  - `queue.go`：队列基础实现。

- **rbtree**：红黑树相关实现（部分功能还未实现）。

  - `rbtree.go`：红黑树基础实现。

- **stack**：栈相关实现。

  - `stack.go`：栈基础实现。

- **vector**：动态数组相关实现。
  - `vector.go`：动态数组基础实现。

# 使用方法

```bash
# 克隆项目到本地
git clone https://github.com/dreamy-xay/golang-data-structure.git
# 或者 interface{} 实现
git clone -b go1.17 https://github.com/dreamy-xay/golang-data-structure.git

# 进入项目目录
cd golang-data-structure

# 运行
go run main.go
```

## 测试

```bash
go test ./...
```
