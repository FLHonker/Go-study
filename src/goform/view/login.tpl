<html>
<head>
    <title>go login</title>
</head>
<body>
    <form action="http://localhost:8811/login" method="POST">
        用户名：<input type="text" name="username">
        </br>
        密码：  <input type="password" name="password">
        </br>
        年龄:   <input type="number" name="age">
        </br>
        邮箱    <input type="text"  name="email">
        </br>
        手机    <input type="text"  name="tele">
        </br>
        身份证号 <input type="text"  name="idcard">
        </br></br>
        <!-- 下拉菜单 -->
        水果：
        <select name="fruit">
        <option value="apple">apple</option>
        <option value="pear">pear</option>
        <option value="banane">banane</option>
        </select>
        </br></br>
        <!-- 单选按钮 -->
        性别：
        <input type="radio" name="gender" value="1">男
        <input type="radio" name="gender" value="2">女
        </br></br>
        <!-- 复选框 -->
        <h3>爱好：</h3>
        <input type="checkbox" name="interest" value="football">足球
        <input type="checkbox" name="interest" value="basketball">篮球
        <input type="checkbox" name="interest" value="tennis">网球
        </br></br>
        <!-- 隐藏字段,tokrn, 用于验证表单提交 -->
        <input type="hidden" name="token" value="{{.}}">
        <input type="submit" value="Login">
    </form>
    <script>
    alert("Hello,Frank!");
    </script>
</body>
</html>