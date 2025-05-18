
# Vaultlet

> **Solution for your API key needs!**  
> Built in Go, cloud-ready, easy to integrate. Just clone & self-host.

---

> [!IMPORTANT]  
> 📌 [WorkFlow & Much more](https://app.eraser.io/workspace/WZJMUDnUQLRQWyRQK2G0?origin=share)

> [!NOTE]  
> ⚙️ Populate the `.env` file with your credentials.

---

### 🚀 Installation

#### Requirements

- [Go 1.24 or above](https://go.dev/)
- [A SQL database](https://letmegooglethat.com/?q=list+of+sql+database)
- [Client application (for testing)](https://http.cat/status/102)

```bash
# Clone into your server or any microservice host
git clone https://github.com/t7abhay/Vaultlet
cd Vaultlet
cp .env_sample .env
```

```bash
# Docker installation
docker build -t vaultlet-app .
docker run --rm --env-file .env -p 9098:9098 vaultlet-app
```

---

### 📘 Documentation

#### 🔑 API Key Generation

```http
POST /api/v1/gen-apikey HTTP/1.1
Content-Type: application/json

{
  "user_id": "some_user_id",     // UUID required
  "duration": 12                 // Optional, in hours. If not set, API key will not expire
}
```

**Response**
```json
HTTP/1.1 200 OK
{
  "success": true,
  "message": "API key created",
  "apiKey": "your_generated_api_key"
}
```

---

#### ✅ API Key Validation

```http
POST /api/v1/validate-apikey HTTP/1.1
Content-Type: application/json

{
  "user_id": "some_user_id",     // UUID required
  "api_key": "your_api_key"
}
```

**Response**
```json
HTTP/1.1 200 OK
{
  "success": true,
  "message": "Valid"
}
```

---

### 📊 Response Code Summary

<table>
  <tr>
    <th>Code</th>
    <th>Meaning</th>
  </tr>
  <tr>
    <td><span style="color:#04d651;"><strong>201</strong></span></td>
    <td>API key created successfully</td>
  </tr>
  <tr>
    <td><span style="color:#04d651;"><strong>200</strong></span></td>
    <td>API key validated</td>
  </tr>
  <tr>
    <td><span style="color:red;"><strong>500</strong></span></td>
    <td>Service failed to validate/create API key</td>
  </tr>
  <tr>
    <td><span style="color:orange;"><strong>401</strong></span></td>
    <td>Invalid API key</td>
  </tr>
  <tr>
    <td><span style="color:orange;"><strong>402</strong></span></td>
    <td>No API key found</td>
  </tr>
  <tr>
    <td><span style="color:goldenrod;"><strong>400</strong></span></td>
    <td>Invalid/Malformed request (Bad Request)</td>
  </tr>
</table>
