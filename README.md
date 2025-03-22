<p align="center">  
  <img src="https://github.com/ROOTUSX/Cloudflare-Checker/blob/main/Screenshot%202025-03-22%20102456.png" alt="Cloudflare Checker" width="600">  
</p>  <h1 align="center">🌐 Cloudflare Checker 🔥</h1>  
<p align="center">  
  <b>A high-performance tool to detect Cloudflare protection, fetch real IPs, and analyze response details.</b>  
  <br>  
  <br>  
  <img src="https://img.shields.io/github/stars/YourUsername/cf-checker?color=yellow&style=flat-square">  
  <img src="https://img.shields.io/github/forks/YourUsername/cf-checker?color=blue&style=flat-square">  
  <img src="https://img.shields.io/github/issues/YourUsername/cf-checker?color=red&style=flat-square">  
  <img src="https://img.shields.io/github/license/YourUsername/cf-checker?color=green&style=flat-square">  
</p>  

🚀 Features

✔ Cloudflare Detection (Headers, JS Challenge, CAPTCHA, HTML Analysis)
✔ Real IP Extraction (if not masked)
✔ Colorized & Structured Output
✔ Status Code Highlighting (Green 🟢, Yellow 🟡, Red 🔴)
✔ Multi-Threaded for Speed
✔ Save Results to a File (-o output.txt)
✔ Animated Startup with Your Name


⚡ Installation & Usage

1️⃣ Install Go (If Not Installed)

sudo apt update && sudo apt install golang -y

2️⃣ Clone & Build

git clone https://github.com/YourUsername/cf-checker.git  
cd cf-checker  
go build -o cf-ck

3️⃣ Run the Tool

 ./cf-ck websites.txt

📌 Save Output to a File:

./cf-ck websites.txt -o results.txt


🎨 Example Output

🔵 Starting Cloudflare Protection Check...  

✅ example.com → Protected by Cloudflare 🟢 [Status: 200] (Time: 123ms)  
   CF-RAY: abc123, CF-Cache: HIT  
   🌍 Real IP: Not Found  

❌ testsite.com → NOT protected by Cloudflare 🔴 [Status: 404] (Time: 78ms)  
   🌍 Real IP: 192.168.1.1, 45.33.21.67  

⚠ failedsite.com → [Error] Connection Timeout


🛠 Advanced Usage

💡 Scan a Single Website

echo "example.com" > single.txt  
./cf-ck single.txt

💡 Increase Timeout (Default: 7s)

./cf-ck websites.txt --timeout=10

💡 Check Websites in a File & Save Results

./cf-ck websites.txt -o cloudflare_results.txt


💻 Contributing

🔥 Want to improve this tool? Fork it, modify, and send a pull request!

git clone https://github.com/YourUsername/cf-checker.git  
cd cf-checker


⚠ Legal Disclaimer

This tool is for educational and research purposes only. Use it responsibly.
The developer is not responsible for any misuse or illegal activities.


<p align="center">  
  🚀 Developed by <b>RootUserX</b> | 🔥 Follow for More!  
</p>  


💎 This README is professionally formatted with badges, a banner, sections, and advanced usage tips.

📌 Replace YourUsername and YourName before uploading.

📢 Want an even fancier version? Add an asciinema GIF of the tool in action! 🚀
