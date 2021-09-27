<template>
  <div class="home">
    <NavbarVue />
    <div class="flex justify-center">
      <!-- modal <add contact> -->
      <label for="my-modal-2" class="btn btn-accent m-2 modal-button"
        >Add Contact</label
      >
      <input type="checkbox" id="my-modal-2" class="modal-toggle" />
      <div class="modal">
        <div class="modal-box">
          <div class="form-control">
            <label class="label">
              <span class="label-text">Name of the contact:</span>
            </label>
            <input
              type="text"
              placeholder="name"
              class="input input-primary input-bordered"
            />
            <label class="label">
              <span class="label-text">Number:</span>
            </label>
            <input
              type="text"
              placeholder="number"
              class="input input-primary input-bordered"
            />
          </div>
          <div class="modal-action">
            <label for="my-modal-2" class="btn btn-success">Save</label>
            <label for="my-modal-2" class="btn btn-ghost">Close</label>
          </div>
        </div>
      </div>
      <!-- end -->
      <AddChat />
    </div>
    <div class="container mx-auto">
      <ChatComponent />
    </div>
  </div>
</template>

<script>
import NavbarVue from '../components/Navbar.vue';
import ChatComponent from '../components/ChatComponent.vue';
import AddChat from '../components/AddChat.vue';
import axios from '../axios';

export default {
  name: 'Home',
  components: {
    NavbarVue,
    ChatComponent,
    AddChat,
  },
  async created() {
    const token = sessionStorage.getItem('token');
    if (!token) {
      this.$router.push('/login');
    }
    const res = await axios.get('/user', {
      headers: { Authorization: `Bearer ${token}` },
    });

    console.log(res);
  },
};
</script>
