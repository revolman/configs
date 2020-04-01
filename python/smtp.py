import smtplib, ssl

mail_server = '89.175.253.238'
mail_server_port = 465

from_addr = 'chulichkov@roadtec.ru'
to_addr = 'admin@roadtec.ru'

from_header = "FROM: %s\r\n" % from_addr
to_header = "TO: %s\r\n\r\n" % to_addr
subject_header = "Subject: Sent by python"

body = "Boring message"

email_message = "%s\n%s\n%s\n\n%s" % (from_header,
                                      to_header,
                                      subject_header,
                                      body)

context = ssl.create_default_context()


s = smtplib.SMTP(mail_server, mail_server_port)
s.ehlo()
s.starttls(context=context)
s.ehlo()
s.login('chulichkov@roadtec.ru', 'PASSWD')
s.sendmail(from_addr, to_addr, email_message)
s.quit()