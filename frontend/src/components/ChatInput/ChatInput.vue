<template>
  <div class="container">
    <input type="text" v-model="message" placeholder="Type your message..." />
    <button @click="sendMessage"><i class="fas fa-paper-plane"></i></button>
  </div>
</template>

<script>
import axios from "../../axios";
export default {
  data() {
    return {
      message: "",
    };
  },
  props: {
    userTo: {
      default: "",
    },
    userFrom: {
      default: "",
    },
  },
  methods: {
    async sendMessage() {
      const user = JSON.parse(sessionStorage.getItem("user"));
      try {
        const res = await axios.post(
          "/messages",
          {
            to: this.userTo,
            from: this.userFrom,
            message: this.message,
          },
          { headers: { Authorization: `Bearer ${user.token}` } }
        );
        this.message = "";
        // location.reload();
      } catch (err) {
        console.log(err);
      }
    },
  },
};
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