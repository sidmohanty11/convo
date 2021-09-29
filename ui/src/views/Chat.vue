<template>
  <div>
    <ChatHeader :chatName="chat.name" />
    <ChatBody :messages="chat.messages" />
    <ChatInput :chatId="chatId" />
  </div>
</template>

<script>
import ChatHeader from '../components/ChatHeader.vue';
import ChatBody from '../components/ChatBody.vue';
import ChatInput from '../components/ChatInput.vue';
import axios from '../axios';

export default {
  components: {
    ChatHeader,
    ChatBody,
    ChatInput,
  },
  data() {
    return {
      token: '',
      chatId: '',
      chat: {
        name: '',
        messages: [],
      },
    };
  },
  async created() {
    this.token = sessionStorage.getItem('token');
    this.chatId = this.$route.params.id;
    if (!this.token) {
      return this.$router.push('/login');
    }
    const { data } = await axios.get(`/chats/${this.chatId}`, {
      headers: { Authorization: `Bearer ${this.token}` },
    });

    this.chat.name = data.name;
    this.chat.messages = data.messages;

    return this.chat;
  },
};
</script>

<style>
</style>
