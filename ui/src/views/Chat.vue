<template>
  <div>
    <ChatHeader />
    <ChatBody />
    <ChatInput />
  </div>
</template>

<script>
import ChatHeader from '../components/ChatHeader.vue';
import ChatBody from '../components/ChatBody.vue';
import ChatInput from '../components/ChatInput.vue';

export default {
  components: {
    ChatHeader,
    ChatBody,
    ChatInput,
  },
};
</script>

<style>
</style>
