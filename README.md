Cloudflare Checker 🔥

Cloudflare Checker is a fast and efficient tool to check if a website is protected by Cloudflare. It also retrieves the real IP address (if available) and displays response details in a colorful output.

Features 🚀

✅ Detects Cloudflare Protection
✅ Finds Real IP Address (if not behind Cloudflare)
✅ Displays HTTP Status Code (Color-Coded)
✅ Shows Response Time
✅ Cloudflare Detection Methods (JS Challenge, CAPTCHA)
✅ Saves Output to a File (-o output.txt option)
✅ Animated Author Name

Installation & Usage 🛠

1. Install Go (if not installed)

sudo apt update && sudo apt install golang -y

2. Clone the Repository

git clone https://github.com/YourUsername/cf-checker.git
cd cf-checker

3. Build the Tool

go build -o cf-ck

4. Run the Tool

./cf-ck websites.txt

Save Output to a File:

./cf-ck websites.txt -o result.txt

Example Output 🎨

✅ example.com → Protected by Cloudflare [Status: 200] (Time: 123ms) 
CF-RAY: xyz123, CF-Cache: HIT [Real IP: Not Found]

Contributing

Pull requests are welcome! If you find a bug or have a feature request, open an issue.

License

This project is for educational purposes only. Use responsibly.


---

🔗 Replace YourUsername with your actual GitHub username before uploading.

Let me know if you need more changes! 
