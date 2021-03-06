### 3 detections

#### ⚠️ Proof Of Concept && Educational Purpose Only ⚠️

## Table of Content
- [Login Showcase](https://github.com/JuPaLe/Token-Grabber#login)
- [Password Changed Showcase](https://github.com/Token-Grabber#changedpassword)
- [Key Features](https://github.com/Token-Grabber#features)
- [Configuration](https://github.com/Token-Grabber#configuration)

### Login
![Login](https://media.discordapp.net/attachments/870608841623085100/901462244527861800/unknown.png?width=454&height=616)
### ChangedPassword
![Password](https://media.discordapp.net/attachments/870608841623085100/901462254875193414/unknown.png?width=517&height=616)

# Features
- Log all friends owning a "rare" badge
- Discord Credit Card Stealing
    - Credit Card Number
    - CVC
    - Credit Card Expirations (month & years)
    - Country
    - State
    - City
    - ZIP Code
    - Address
    - Username
    - Nitro
    - Badges
    - ID
    - Email
    - Token
- Discord Login Stealing
    - Username
    - Nitro
    - Badges
    - ID
    - Email
    - Password
    - Token
    - Valid Billing and type
- Discord Change Password
    - Username
    - Nitro
    - Badges
    - ID
    - Old Password
    - New Password
    - Token
    - Valid Billing and type
- Discord Change Email
    - Username
    - Nitro
    - Badges
    - ID
    - New Email
    - Password
    - Token
    - Valid Billing and type
- IP
- Remove QR Code
- BetterDiscord Protection Remover

# Configuration
|            Name | Description                                                                                                                                          | Example     | Type    |
|----------------:|------------------------------------------------------------------------------------------------------------------------------------------------------|-------------|---------|
| platform        | An array containing all the platforms you want your grabber to work on.                                                                              | ["windows"] | array   |
| logout          | False: Do not log out delayed: Disconnect the next time you launch Discord Instant: Kill Discord and ask to reconnect                                   | "instant"   | boolean |
| inject-notify   | False: Does not send a message when the grabber has successfully injected True: Send a message when the grabber has successfully injected            | "true"      | boolean |
| logout-notify   | False: Does not send a message when the victim has successfully logged out True: Send a message when the victim has successfully logged out          | "false"     | boolean |
| init-notify     | False: Does not send a message when the grabber has been initialized in the victim discord client True: Send the message                             | "false"     | boolean |
| embed-color     | The embed color in decimal!                                                                                                                              | "3447704"   | string  |
| disable-qr-code | False: The victim will have access to the authentication QRCode (not recommended) True: The victim will not have access to the authentication QRCode | "true"      | boolean |
