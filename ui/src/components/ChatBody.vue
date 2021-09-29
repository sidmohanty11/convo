<template>
  <div class="bodyOfChat container mx-auto">
    <div
      class="card p-2 my-2 shadow-md"
      v-for="message in messages"
      :key="message.id"
    >
      <p>
        <span class="font-bold"
          >{{ message.username }}
          <span class="text-green-500 text-2xs"
            >({{ new Date(message.timestamp).toLocaleString() }})</span
          >
          : </span
        >{{ message.content }}
      </p>
    </div>
  </div>
</template>

<script>
export default {
  props: ['messages'],
};
</script>

<style>
.bodyOfChat {
  height: calc(100vh - 15rem);
}
</style>
