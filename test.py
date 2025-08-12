import requests

url = "http://whatsapp.findanime.to/api/sendText"
headers = {
    "Accept": "application/json",
    "Content-Type": "application/json"
}
data = {
    "chatId": "918910114007@s.whatsapp.net",  # Replace with the actual chat ID
    "text": """
🎉 *Welcome to Club 0day!* 🎉

You're officially a member! 🚀  
We're thrilled to have you with us. Here's how you can stay connected:

💬 *Join our community*:  
  - *WhatsApp*: https://chat.whatsapp.com/IMRoMajx2J27FNoPMi5VzC  
  - *Discord*: https://discord.gg/VB4FMbuukE

📞 *Need help?* Reach out to Arya: 7603061337

Let's innovate, learn, and grow together! 💡  
Excited to see you in action! 🔥
    """,
    "session": "default"
}

response = requests.post(url, json=data, headers=headers)
print(response.json())
