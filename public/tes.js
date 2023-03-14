// make the request to the login endpoint
async function getToken() {
  username = document.getElementById("username").value;
  password = document.getElementById("password").value;
  body = {
    username: username,
    password: password,
  };
  const res = await fetch("/api/login", {
    method: "POST",
    headers: {
      "content-type": "application/json",
    },
    redirect: "follow",
    referrerPolicy: "no-referrer",
    body: JSON.stringify(body),
  });
  token = (await res.json()).token
  localStorage.setItem("token", "Bearer " + token)
  tokenElement = document.getElementById('token')
  tokenElement.innerHTML = token
}

async function getSecret() {
  token = localStorage.getItem('token')
  const res = await fetch("/api/menu", {
    method: "GET",
    headers: {
      "Authorization": token
    },
    redirect: "follow", 
    referrerPolicy: "no-referrer"
  });
  secretElement = document.getElementById('result')
  secretElement.innerHTML = await res.json()
}
