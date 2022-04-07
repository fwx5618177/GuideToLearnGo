# 数据库设计
1. User
2. Session
3. Thread
4. Post

# 绕过静态类型
- 任何类型的值都可以当作空接口类型: `func generateHTML(w http.ResponseWriter, data interface{}, fn ...string) {}`