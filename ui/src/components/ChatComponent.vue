<template>
  <div class="card lg:card-side bordered shadow-md m-3">
    <div class="card-body">
      <h2 class="card-title text-center">{{ chatName }}</h2>
      <button class="btn btn-primary">Chat</button>
    </div>
  </div>
</template>

<script>
export default {
  props: ['chatName'],
};
</script>

<style>
</style>
