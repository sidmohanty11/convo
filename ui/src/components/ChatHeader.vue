<template>
  <div
    class="navbar mb-2 shadow-lg bg-neutral text-neutral-content rounded-box"
  >
    <div class="flex-none">
      <button @click="$router.push('/')" class="btn btn-square btn-ghost">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          class="inline-block w-6 h-6 stroke-current text-success"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M10 19l-7-7m0 0l7-7m-7 7h18"
          ></path>
        </svg>
      </button>
    </div>
    <div class="px-2 mx-2">
      <h1 class="text-lg font-bold">#{{ chatName }}</h1>
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
