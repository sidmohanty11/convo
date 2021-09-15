<template>
  <div class="home">
    <router-link to="/login"
      ><i class="fas fa-chevron-left"
        ><span class="span"> Back</span></i
      ></router-link
    >
    <h1 style="font-size: 28px">
      Welcome to <span class="span">Convo!</span> @{{ user?.username }}
    </h1>
    <div class="search">
      <input class="search-input" placeholder="Search" />
    </div>
    <ChatComponent
      v-for="chat in chats"
      :key="chat.id"
      :chat="chat"
      :pushFunc="pushFunc"
    />

    <div class="add-chat">
      <button @click="addChat" class="add-chat-icon">
        <i class="fas fa-plus"></i>
      </button>
    </div>
  </div>
</template>

<script>
import ChatComponent from "../components/ChatComponent/ChatComponent.vue";
import axios from "../axios";
export default {
  name: "Home",
  components: {
    ChatComponent,
  },
  data() {
    return {
      user: {},
      chats: [],
    };
  },
  methods: {
    pushFunc(id) {
      this.$router.push(`/message/${id}`);
    },
    async addChat() {
      const username = prompt("Who do you want to chat with?");
      if (username) {
        const res = await axios.post(
          "/messages",
          {
            to: username,
            message: "Hello!",
            from: this.user.username,
          },
          {
            headers: { Authorization: `Bearer ${this.user.token}` },
          }
        );
        console.log(res);
        if (res) {
          this.pushFunc(res.data.to);
        }
      } else {
        console.log("error bro");
      }
    },
  },
  async created() {
    const user = JSON.parse(sessionStorage.getItem("user"));
    if (user && user.token) {
      this.user = user;
      const res = await axios.get(`/messages/${user.username}`, {
        headers: { Authorization: `Bearer ${user.token}` },
      });
      if (res) {
        this.chats = res.data;
      }
      console.log(res);
    } else {
      this.$router.push("/login");
    }
  },
};
</script>

<style scoped>
.home {
  margin: 10px;
  padding: 10px;
}

.addChat {
  margin: 0 auto;
}

.add-chat-icon {
  position: absolute;
  padding: 20px;
  margin: 10px;
  border-radius: 100%;
  border: none;
  top: 85%;
  left: 80%;
  background-color: #b10e0b;
  cursor: pointer;
}

.search {
  display: grid;
  place-items: center;
}

.search-input {
  width: 80%;
  height: 10px;
  padding: 15px;
  margin: 5px;
  border-radius: 6px;
}

i {
  color: whitesmoke;
}

.span {
  background: #fff;
  background: -webkit-linear-gradient(to right, #fff 0%, #cf1512 100%);
  background: -moz-linear-gradient(to right, #fff 0%, #cf1512 100%);
  background: linear-gradient(to right, #fff 0%, #cf1512 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}
</style>