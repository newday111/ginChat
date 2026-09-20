package emailformat

import "fmt"

func RegisterCodeFormat(code string) string {
	body := fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<head>
			<meta charset="UTF-8">
			<title>注册验证码</title>
		</head>
		<body style="margin: 0; padding: 0; background-color: #f5f7fa; font-family: Arial, sans-serif;">

			<div style="max-width: 600px; margin: 40px auto; padding: 0 20px;">

				<div style="background-color: #ffffff; border-radius: 10px; padding: 40px;">

					<h2 style="margin: 0 0 20px; color: #333333;">
						注册验证码
					</h2>

					<p style="font-size: 16px; color: #555555;">
						您好！
					</p>

					<p style="font-size: 16px; color: #555555;">
						您正在进行账号注册，本次验证码为：
					</p>

					<div style="margin: 30px 0; text-align: center;">
						<span style="
							display: inline-block;
							padding: 15px 30px;
							background-color: #f0f5ff;
							border-radius: 8px;
							font-size: 32px;
							font-weight: bold;
							letter-spacing: 8px;
							color: #1677ff;
						">
							%s
						</span>
					</div>

					<p style="font-size: 14px; color: #999999;">
						验证码 <strong>1 分钟</strong>内有效，请勿将验证码泄露给其他人。
					</p>

					<p style="font-size: 14px; color: #999999;">
						如果不是您本人操作，请忽略此邮件。
					</p>

					<hr style="border: 0; border-top: 1px solid #eeeeee; margin: 30px 0;">

					<p style="font-size: 12px; color: #bbbbbb; text-align: center;">
						此邮件由系统自动发送，请勿直接回复。
					</p>

				</div>

			</div>

		</body>
		</html>
		`, code)
	return body
}
