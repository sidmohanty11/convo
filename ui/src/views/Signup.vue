<template>
  <div>
    <div
      class="
        flex
        items-center
        w-full
        h-screen
        px-4
        py-10
        bg-cover
        card
        bg-base-200
      "
      style="
        background-image: url('https://images.unsplash.com/photo-1607357221935-9133bb4b44ca?ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&ixlib=rb-1.2.1&auto=format&fit=crop&w=1470&q=80');
      "
    >
      <div class="card glass lg:card-side text-neutral-content">
        <figure class="p-6">
          <img
            src="https://images.unsplash.com/photo-1512626120412-faf41adb4874?ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&ixlib=rb-1.2.1&auto=format&fit=crop&w=300&q=80"
            class="rounded-lg shadow-lg"
          />
        </figure>
        <div class="max-w-md card-body">
          <h2 class="card-title text-gray-400 text-center">Signup | Convo.</h2>
          <div class="form-control">
            <input
              type="text"
              placeholder="number"
              class="input input-bordered m-2"
            />
            <input
              type="text"
              placeholder="password"
              class="input input-bordered m-2"
            />
          </div>
          <div class="card-actions">
            <p class="text-gray-200">
              Already have an account?
              <router-link class="text-gray-800" to="/login"
                >Login!</router-link
              >
            </p>
          </div>
          <button class="btn glass rounded-full">Login</button>
        </div>
      </div>
    </div>
    <footer class="p-4 footer bg-base-300 text-base-content footer-center">
      <div>
        <p>Copyright © 2021 - All right reserved by Sidharth Ltd</p>
      </div>
    </footer>
  </div>
</template>

<script>
export default {};
</script>

<style>
</style>
