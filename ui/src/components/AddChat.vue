<template>
  <label for="addChat" class="btn btn-accent m-2 modal-button">Add Chat</label>
  <input type="checkbox" id="addChat" class="modal-toggle" />
  <div class="modal">
    <div class="modal-box">
      <span class="label-text text-red-500">{{ err }}</span>
      <select
        v-model="selectedOption"
        class="select select-bordered select-primary w-full max-w-xs"
      >
        <option value="group">Group Chat</option>
        <option value="dm">Personal Chat</option>
      </select>
      <!-- group -->
      <div v-if="selectedOption === 'group'" class="form-control">
        <label class="label">
          <span class="label-text">Name of the Group:</span>
        </label>
        <input
          v-model="addChatInfo.name"
          type="text"
          placeholder="name of group"
          class="input input-primary input-bordered"
        />
        <p class="mt-2">Contacts:</p>
        <div
          class="mt-2 card bordered"
          v-for="contact in contacts"
          :key="contact._id"
        >
          <label class="cursor-pointer label">
            <span class="label-text">{{ contact.saved_as }}</span>
            <input
              type="checkbox"
              class="checkbox checkbox-accent"
              :id="contact._id"
              v-model="addChatInfo.participants"
              :value="contact._id"
            />
          </label>
        </div>
      </div>
      <!-- personal -->
      <div v-else>
        <div class="form-control mt-2">
          <label
            class="cursor-pointer label"
            v-for="contact in contacts"
            :key="contact._id"
          >
            <span class="label-text">{{ contact.saved_as }}</span>
            <input
              v-model="selectedUserForDM"
              type="radio"
              name="opt"
              class="radio radio-accent"
              :value="contact"
            />
          </label>
        </div>
      </div>
      <div class="modal-action">
        <label @click="addChat" class="btn btn-success">Create</label>
        <label for="addChat" class="btn btn-ghost">Close</label>
      </div>
    </div>
  </div>
</template>

<script>
import axios from '../axios';

export default {
  data() {
    return {
      selectedOption: 'group',
      addChatInfo: {
        name: '',
        participants: [],
      },
      selectedUserForDM: {},
      err: '',
    };
  },
  props: ['contacts', 'userId'],
  methods: {
    async addChat() {
      const token = sessionStorage.getItem('token');
      if (this.selectedOption === 'group' && this.addChatInfo.name !== '') {
        this.addChatInfo.participants.push(this.userId);

        const res = await axios.post(
          '/chat/add',
          {
            ...this.addChatInfo,
          },
          {
            headers: { Authorization: `Bearer ${token}` },
          },
        );

        console.log(res);
      } else if (this.selectedOption === 'dm') {
        // eslint-disable-next-line no-underscore-dangle
        const participants = [this.selectedUserForDM._id, this.userId];
        const name = this.selectedUserForDM.saved_as;

        if (participants.length < 2 && name === '') {
          this.err = 'Do not even try to fool me bro!';
          return this.err;
        }

        const res = await axios.post(
          '/chat/add',
          {
            name,
            participants,
          },
          {
            headers: { Authorization: `Bearer ${token}` },
          },
        );

        console.log(res);
      }

      this.err = 'Make sure to provide correct details!';
      return this.err;
    },
  },
};
</script>

<style>
</style>
