<template>
  <div class="home">
    <NavbarVue :username="user.username" :contacts="user.contacts" />
    <p class="text-center">
      Your last login was at:
      <span class="text-green-500 font-bold">{{
        new Date(user.last_seen).toLocaleString()
      }}</span>
    </p>
    <div class="flex justify-center">
      <!-- modal <add contact> -->
      <label for="my-modal-2" class="btn btn-accent m-2 modal-button"
        >Add Contact</label
      >
      <input type="checkbox" id="my-modal-2" class="modal-toggle" />
      <div class="modal">
        <div class="modal-box">
          <span class="label-text text-red-500">{{ err }}</span>
          <div class="form-control">
            <label class="label">
              <span class="label-text">Name of the contact:</span>
            </label>
            <input
              v-model="addContactInfo.saved_as"
              type="text"
              placeholder="name"
              class="input input-primary input-bordered"
            />
            <label class="label">
              <span class="label-text">Number:</span>
            </label>
            <input
              v-model="addContactInfo.number"
              type="text"
              placeholder="number"
              class="input input-primary input-bordered"
            />
          </div>
          <div class="modal-action">
            <label class="btn btn-success" @click="addContact">Save</label>
            <label for="my-modal-2" class="btn btn-ghost">Close</label>
          </div>
        </div>
      </div>
      <!-- end -->
      <AddChat :contacts="user.contacts" :userNumber="user.number" />
    </div>
    <div class="container mx-auto">
      <ChatComponent
        v-for="chat in user.chats"
        :key="chat.id"
        @click="$router.push(chat.id)"
      />
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
  data() {
    return {
      user: {
        chats: [],
        contacts: [],
        id: '',
        last_seen: '',
        number: '',
        username: '',
      },
      addContactInfo: {
        saved_as: '',
        number: '',
      },
      token: '',
      err: '',
    };
  },
  components: {
    NavbarVue,
    ChatComponent,
    AddChat,
  },
  async created() {
    this.token = sessionStorage.getItem('token');
    if (!this.token) {
      this.$router.push('/login');
    }
    const { data } = await axios.get('/activeUser', {
      headers: { Authorization: `Bearer ${this.token}` },
    });

    this.user = data;
  },
  methods: {
    async addContact() {
      if (this.addContactInfo.saved_as === '') {
        this.err = 'You can not fool me bro!';
      } else if (this.addContactInfo.number.length !== 10) {
        this.err = 'Numbers in India are of 10 digits, right?';
      }

      if (this.err !== '') {
        return this.err;
      }

      await axios.post(
        '/contacts/add',
        {
          ...this.addContactInfo,
        },
        {
          headers: { Authorization: `Bearer ${this.token}` },
        },
      );

      return window.location.reload();
    },
  },
};
</script>
