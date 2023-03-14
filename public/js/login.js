async function onLoginClick() {
  let username = document.getElementById("username").value;
  let password = document.getElementById("password").value;
  let body = {
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
  let token = (await res.json()).token
  localStorage.setItem('token', 'Bearer ' + token)
}
