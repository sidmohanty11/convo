<template>
  <div class="navbar mb-2 shadow-lg bg-neutral text-neutral-content">
    <div class="flex-1 px-2 mx-2">
      <span class="text-lg font-bold"> Convo </span>
    </div>
    <div class="flex-none">
      <button class="btn btn-ghost">hi #NAME</button>
      <label for="my-modal" class="btn btn-ghost modal-button"
        >My Contacts</label
      >
      <input type="checkbox" id="my-modal" class="modal-toggle" />
      <div class="modal">
        <div class="modal-box">
          <contacts />
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
};
</script>

<style>
</style>
