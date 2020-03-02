AWS Amazon Linux 2 AMI + Golang + SSL Letsencrypt
=======

From 0 to AMI2 + Go/Golang + SSL in 15 min or less

Description
=======

	Setup a web in 15 min or less

## Installation

	// Prerequisites
		- A domain -anywhere- (NOTE: remove AAAA DNS)
			* open port 80, Letsencrypt need it.
			* open 8080 and 443 ports, your app need it.
		- An Amazon AMI2 instance

	// Perform a quick software update on your instance
		sudo yum update -y

	// Letsencrypt Install
		sudo yum install python27-devel git -y
		sudo git clone https://github.com/letsencrypt/letsencrypt /opt/letsencrypt

	// Edit the file /opt/letsencrypt/certbot-auto to recognize your version of Linux
		sudo nano /opt/letsencrypt/certbot-auto
	// find this line in the file (likely near line nr 780):
		elif [ -f /etc/redhat-release ]; then
	// and replace whole line with this:
		elif [ -f /etc/redhat-release ] || grep 'cpe:.*:amazon_linux:2' /etc/os-release > /dev/null 2>&1; then

	// Run the letsencrypt setup, and follow the instructions
		sudo /opt/letsencrypt/certbot-auto certonly -d YOU_DOMAIN -d www.YOU_DOMAIN --manual --preferred-challenges dns-01 --force-renewal --manual-public-ip-logging-ok

	// Upload the compiled file and the application service file
		sudo chmod 755 application	
		sudo chmod 755 application.service

	// Modify access permisions to the certs
		sudo chmod 755 /etc/letsencrypt/live/
		sudo chmod 755 /etc/letsencrypt/archive/
		sudo chmod 755 /etc/letsencrypt/live/YOU_DOMAIN
		sudo chmod 755 /etc/letsencrypt/archive/YOU_DOMAIN

	// Forwarding the traffic from 80 to 8080 port
		sudo iptables -t nat -I PREROUTING -p tcp --dport 80 -j REDIRECT --to-ports 8080
		sudo iptables -t nat -I OUTPUT -p tcp -o lo --dport 80 -j REDIRECT --to-ports 8080

	// Once uploaded the compiled file and the application service file
	// Start the service
		- sudo mv application.service /etc/systemd/system && sudo systemctl daemon-reload && sudo systemctl enable application

	// Manage your application
		- sudo setcap CAP_NET_BIND_SERVICE=+eip /home/ec2-user/golang/src/PROJECT_DIR/application && sudo systemctl restart application
		- sudo systemctl start application
		- sudo systemctl restart application

	// Recomended after all the steps
		*** restart the instance

	// LOGS
		sudo systemctl status application -l
		journalctl -u application.service --no-page

## References

// LETSENCRYPT ON AMAZON AMI2
- https://medium.com/@gnowland/deploying-lets-encrypt-on-an-amazon-linux-ami-ec2-instance-f8e2e8f4fc1f
- https://medium.com/@andrenakkurt/great-guide-thanks-for-putting-this-together-gifford-nowland-c3ce0ea2455

// FORWARDING RULE PORT TO ANOTHER PORT
- https://googlecloudplatform.uservoice.com/forums/302595-compute-engine/suggestions/8518255-need-to-route-from-forwarding-rule-port-to-another

// SETCAP ISSUE, OPENING PORT 443
- https://stackoverflow.com/questions/54690289/golang-on-gcp-listen-tcp-443-bind-permission-denied
- https://unix.stackexchange.com/questions/455221/setcap-not-found-in-debian-9/455234#455234
