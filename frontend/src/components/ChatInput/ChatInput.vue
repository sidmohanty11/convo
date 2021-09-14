<template>
  <div class="container">
    <input type="text" placeholder="Type your message..." />
    <button><i class="fas fa-paper-plane"></i></button>
  </div>
</template>

<script>
export default {};
</script>

<style scoped>
.container {
  position: absolute;
  top: 87%;
}

input {
  width: 70%;
  height: 10px;
  padding: 15px;
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
</style>