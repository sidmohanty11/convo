<template>
  <div class="login">
    <h1>CONVO.</h1>
    <h2>Login</h2>
    <input type="email" placeholder="email" />
    <input type="password" placeholder="password" />
    <p>
      Don't have an account yet? <router-link to="/signup">Signup.</router-link>
    </p>
    <button><i class="fas fa-arrow-right"></i></button>
  </div>
</template>

<style scoped>
.login {
  display: grid;
  place-items: center;
}
h1 {
  font-size: 60px;
  background: #fff;
  background: -webkit-linear-gradient(to right, #fff 0%, #cf1512 100%);
  background: -moz-linear-gradient(to right, #fff 0%, #cf1512 100%);
  background: linear-gradient(to right, #fff 0%, #cf1512 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}
h2 {
  font-size: 40px;
}
input {
  width: 80%;
  height: 10px;
  padding: 15px;
  margin: 5px;
  border-radius: 6px;
}
button {
  padding: 20px;
  margin: 10px;
  border-radius: 100%;
  border: none;
  color: whitesmoke;
  background-color: #b10e0b;
  cursor: pointer;
}
a {
  color: #b10e0b;
  font-weight: 600;
}
</style>