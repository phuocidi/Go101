<html>
<head>
	<title>Upload File</title>
</head>
<body>
	<form action="http://127.0.0.1:9090/upload" enctype="multipart/form-data" method="post">
		<input type="file" name="uploadfile">
		<input type="hidden" name="token" value="{{.}}">
		<input type="submit" value="upload">
	</form>
</body>
</html>