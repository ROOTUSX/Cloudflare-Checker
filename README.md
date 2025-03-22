🌐 Cloudflare Checker 🔥

A fast, efficient, and colorful tool to detect Cloudflare protection, fetch real IP addresses, and display detailed response information. Perfect for security research and network analysis!


---

🚀 Features

✅ Detects Cloudflare Protection in seconds
✅ Retrieves Real IP Address (if not hidden behind Cloudflare)
✅ Color-Coded HTTP Status Codes for better readability
✅ Displays Response Time for performance analysis
✅ Cloudflare Detection Methods (JS Challenge, CAPTCHA, Headers, HTML Body)
✅ Animated Author Name for a cool startup effect
✅ Saves Output to a File using the -o output.txt flag
✅ Multi-Threaded for fast scanning


---

🛠 Installation & Usage

1️⃣ Install Go (if not installed)

sudo apt update && sudo apt install golang -y

2️⃣ Clone the Repository

git clone https://github.com/YourUsername/cf-checker.git
cd cf-checker

3️⃣ Build the Tool

go build -o cf-ck

4️⃣ Run the Tool

./cf-ck websites.txt

💾 Save Output to a File:

./cf-ck websites.txt -o results.txt


---

🎨 Example Output

🔵 Starting Cloudflare Protection Check...

✅ example.com → Protected by Cloudflare 🟢 [Status: 200] (Time: 123ms)
   CF-RAY: abc123, CF-Cache: HIT
   🌍 Real IP: Not Found

❌ testsite.com → NOT protected by Cloudflare 🔴 [Status: 404] (Time: 78ms)
   🌍 Real IP: 192.168.1.1, 45.33.21.67

⚠ failedsite.com → [Error] Connection Timeout


---

👨‍💻 Contributing

Pull requests are welcome! If you find a bug or have a feature request, open an issue.


---

⚠ Legal Disclaimer

This tool is intended for educational purposes only. Use responsibly. The author is not responsible for any misuse.


---

📌 Replace YourUsername with your actual GitHub username before uploading.

💎 This README is optimized for clarity, beauty, and professionalism! Let me know if you want further tweaks! 🚀
