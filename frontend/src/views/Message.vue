<template>
  <div class="container">
    <MessageHeader />
    <main>
      <div class="chat right">
        <p>hi</p>
        <span class="timestamp">5 mins ago</span>
      </div>
      <div class="chat left">
        <p>hi</p>
        <span class="timestamp">5 mins ago</span>
      </div>
    </main>
    <ChatInput />
  </div>
</template>

<script>
import MessageHeader from "../components/MessageHeader/MessageHeader.vue";
import ChatInput from "../components/ChatInput/ChatInput.vue";
export default {
  data() {
    return {
      paramId: this.$route.params.id,
    };
  },
  components: {
    MessageHeader,
    ChatInput,
  },
};
</script>

<style scoped>
.container {
  padding: 10px;
  margin: 10px;
}

main {
  background-image: url("../assets/bg1.jpg");
  max-height: 1000px;
  min-height: 770px;
  width: 92%;
  overflow-y: scroll;
}

main::-webkit-scrollbar {
  display: none;
}

main {
  -ms-overflow-style: none;
  scrollbar-width: none;
}
.chat {
  width: fit-content;
  padding: 15px;
  border-radius: 8px;
  margin: 10px;
  min-width: 60px;
  padding-bottom: 26px;
  position: relative;
  text-align: right;
}
.right {
  margin-left: auto;
  background-color: #641220;
  color: whitesmoke;
}

.left {
  background-color: black;
  text-align: left;
}

.timestamp {
  color: whitesmoke;
  padding: 10px;
  font-size: 9px;
  position: absolute;
  bottom: 0;
  text-align: right;
  right: 0;
}
</style>