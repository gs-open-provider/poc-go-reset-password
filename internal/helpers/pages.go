package helpers

// IndexPage Constant
const IndexPage = `
<html>
    <head>
			<title>Reset Password</title>
			<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.4.0/css/bootstrap.min.css">
			<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.4.0/jquery.min.js"></script>
			<script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.4.0/js/bootstrap.min.js"></script>
			<script type="text/javascript">

				resetPassword = () => {
					console.log("Reset Password Clicked..");
					$.ajax({
						url: "http://192.168.60.11:9090/reset-password",
						type: 'post',
						data: JSON.stringify({
								username: "user",
								email: "prashanth@gapstars.net"
							}),
						success: function(data) {
							result = JSON.parse(data);
							console.log(result);
						}
					});
				}

			</script>
    </head>
    <body>
			<div class="container">
				<h2>POC on Reset Password in GO</h2>
				<p>
					<button class="btn btn-primary" onclick="resetPassword()">Reset Password</button>
				</p>
			</div>
    </body>
</html>
`

// CallBackPage Constant
const CallBackPage = `
<html>
    <head>
			<title>Reset Password</title>
			<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.4.0/css/bootstrap.min.css">
			<script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.4.0/js/bootstrap.min.js"></script>
    </head>
    <body>
			<div class="container">
				<h2>POC on Reset Password in GO</h2>
				<p>
					<strong>Please reset your password..</strong>
				</p>
			</div>
    </body>
</html>
`
