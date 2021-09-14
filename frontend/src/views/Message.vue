<template>
  <div class="about">
    <h1>This is message page{{ paramId }}</h1>
  </div>
</template>

<script>
export default {
  data() {
    return {
      paramId: this.$route.params.id,
    };
  },
};
</script>