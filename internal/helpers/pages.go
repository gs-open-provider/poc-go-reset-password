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
					console.log("Reset Password Clicked..", $('#email').val());
					if ($('#email').val() !== "") {
						$.ajax({
							url: "http://192.168.60.11:9090/reset-password",
							type: 'post',
							data: JSON.stringify({
									username: $('#username').val(),
									email: $('#email').val()
								}),
							success: function(data) {
								console.log("Email Sent..");
								window.location.href = "http://192.168.60.11:9090/email-sent";
							}
						});
					} else {
						console.log("Please enter email address..");
					}
				}

			</script>
    </head>
    <body>
			<div class="container">
				<h2>POC on Reset Password in GO</h2>
				<div style="padding-top: 25px; padding-left: 100px;">
					<table>
						<tr>
							<td style="width: 100px;"><label>UserName:</label></td>
							<td><label><input id="username" type="text" name="username" /></td>
						</tr>
						<tr>
							<td style="width: 100px;"><label>Email:</label></td>
							<td><label><input id="email" type="text" name="email" /></td>
						</tr>
						<tr>
							<td></td>
							<td><button class="btn btn-primary pull-right" onclick="resetPassword()">Reset Password</button></td>
						</tr>
				</div>
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
