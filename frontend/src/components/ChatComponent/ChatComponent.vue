<template>
  <div class="container">
    <img class="user-image" :src="imageUrl" />
    <button class="chat-component">@{{ username }}</button>
  </div>
</template>

<script>
export default {
  props: {
    username: {
      required: true,
    },
    imageUrl: {
      default: "",
    },
  },
};
</script>

<style scoped>
.container {
  max-height: 57px;
  margin: 10px;
  padding: 10px;
  width: 90%;
  border-radius: 7px;
  display: flex;
  background-color: #222;
}

.chat-component {
  border: none;
  width: 76%;
  font-size: 20px;
  color: whitesmoke;
  background-color: transparent;
}
.user-image {
  max-height: 50px;
  width: 50px;
  border-radius: 100%;
}
</style>