<template>
  <div class="p-10 card bg-base-200">
    <form class="form-control" @submit.prevent="sendMessage">
      <input
        v-model="content"
        type="text"
        placeholder="write your message here..."
        class="input input-ghost bottomInput"
      />
    </form>
    <p class="text-red-500">{{ err }}</p>
  </div>
</template>

<script>
import axios from '../axios';

export default {
  props: ['chatId'],
  data() {
    return {
      content: '',
      username: '',
      err: '',
    };
  },
  methods: {
    async sendMessage() {
      const token = sessionStorage.getItem('token');

      const { data } = await axios.get('/activeUser', {
        headers: { Authorization: `Bearer ${token}` },
      });

      this.username = data.username;

      const { username, content } = this;

      if (username === '' || content === '') {
        this.err = 'you cannot fool me bro.';
        return this.err;
      }

      const { data: message } = await axios.post(
        `/messages/add/${this.chatId}`,
        {
          username,
          content,
        },
        {
          headers: { Authorization: `Bearer ${token}` },
        },
      );

      this.content = '';
      return console.log(message);
    },
  },
};
</script>

<style>
</style>
