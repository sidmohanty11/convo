<template>
  <label for="addChat" class="btn btn-accent m-2 modal-button">Add Chat</label>
  <input type="checkbox" id="addChat" class="modal-toggle" />
  <div class="modal">
    <div class="modal-box">
      <select
        v-model="selectedOption"
        class="select select-bordered select-primary w-full max-w-xs"
      >
        <option value="group">Group Chat</option>
        <option value="dm">Personal Chat</option>
      </select>
      <div v-if="selectedOption === 'group'" class="form-control">
        <label class="label">
          <span class="label-text">Name of the Group:</span>
        </label>
        <input
          type="text"
          placeholder="name of group"
          class="input input-primary input-bordered"
        />
        <p class="mt-2">Contacts:</p>

        <div class="mt-2 card bordered">
          <label class="cursor-pointer label">
            <span class="label-text">names</span>
            <input
              type="checkbox"
              checked="checked"
              class="checkbox checkbox-accent"
            />
          </label>
        </div>
        <div class="mt-2 card bordered">
          <label class="cursor-pointer label">
            <span class="label-text">names</span>
            <input type="checkbox" class="checkbox checkbox-accent" />
          </label>
        </div>
        <div class="mt-2 card bordered">
          <label class="cursor-pointer label">
            <span class="label-text">names</span>
            <input type="checkbox" class="checkbox checkbox-accent" />
          </label>
        </div>
      </div>
      <div v-else>
        <div class="form-control mt-2">
          <label class="cursor-pointer label">
            <span class="label-text">Names</span>
            <input
              type="radio"
              name="opt"
              checked="checked"
              class="radio radio-accent"
              value=""
            />
          </label>
          <label class="cursor-pointer label">
            <span class="label-text">Names</span>
            <input
              type="radio"
              name="opt"
              class="radio radio-accent"
              value=""
            />
          </label>
        </div>
      </div>
      <div class="modal-action">
        <label for="addChat" class="btn btn-success">Create</label>
        <label for="addChat" class="btn btn-ghost">Close</label>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      selectedOption: 'group',
    };
  },
};
</script>

<style>
</style>
