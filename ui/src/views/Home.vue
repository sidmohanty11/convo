<template>
  <div class="home">
    <NavbarVue />
    <div class="flex justify-center">
      <button class="btn btn-primary m-2">Add Contact</button>
      <button class="btn btn-primary m-2">Add Chat</button>
    </div>
    <div class="container mx-auto">
      <ChatComponent />
    </div>
  </div>
</template>

<script>
import NavbarVue from '../components/Navbar.vue';
import ChatComponent from '../components/ChatComponent.vue';

export default {
  name: 'Home',
  components: {
    NavbarVue,
    ChatComponent,
  },
};
</script>
