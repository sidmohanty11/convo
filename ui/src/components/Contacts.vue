<template>
  <div class="overflow-x-auto">
    <table class="table w-full text-black">
      <thead>
        <tr>
          <th>Name</th>
          <th>Contact Number</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="contact in contacts" :key="contact.id">
          <td>{{ contact.saved_as }}</td>
          <td>{{ contact.number }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  props: ['contacts'],
};
</script>

<style>
</style>
