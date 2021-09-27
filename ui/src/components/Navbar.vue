<template>
  <div class="navbar mb-2 shadow-lg bg-neutral text-neutral-content">
    <div class="flex-1 px-2 mx-2">
      <span class="font-bold text-3xl logo-font"> Convo. </span>
    </div>
    <div class="flex-none">
      <button class="btn btn-ghost" @click="logout">
        hi,<span class="text-red-600"> #{{ username }}</span>
      </button>
      <label for="my-modal" class="btn btn-ghost modal-button"
        >My Contacts</label
      >
      <input type="checkbox" id="my-modal" class="modal-toggle" />
      <div class="modal">
        <div class="modal-box">
          <Contacts :contacts="contacts" />
          <div class="modal-action">
            <label for="my-modal" class="btn">Close</label>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Contacts from './Contacts.vue';

export default {
  components: { Contacts },
  props: ['username', 'contacts'],
  methods: {
    logout() {
      sessionStorage.removeItem('token');
      return this.$router.push('/login');
    },
  },
};
</script>

<style>
@import url('https://fonts.googleapis.com/css2?family=Yomogi&display=swap');

.logo-font {
  font-family: 'Yomogi', cursive;
}
</style>
