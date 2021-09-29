<template>
  <div>
    <div class="alert alert-error" v-show="err !== ''">
      <div class="flex-1">
        <label>{{ err }}</label>
      </div>
    </div>
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
      style="background-color: black"
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
              v-model="formData.number"
              type="text"
              placeholder="number"
              class="input input-bordered m-2"
            />
            <input
              v-model="formData.username"
              type="text"
              placeholder="username"
              class="input input-bordered m-2"
            />
            <input
              v-model="formData.password"
              type="password"
              placeholder="password"
              class="input input-bordered m-2"
            />
          </div>
          <div class="card-actions">
            <p class="text-gray-200">
              Already have an account?
              <router-link class="text-gray-400" to="/login"
                >Login!</router-link
              >
            </p>
          </div>
          <button class="btn glass rounded-full text-gray-400" @click="signup">
            Signup
          </button>
        </div>
      </div>
    </div>
    <footer class="p-4 footer text-base-content footer-center bg-black">
      <div>
        <p>Copyright © 2021 - All right reserved by Sidharth Ltd</p>
      </div>
    </footer>
  </div>
</template>

<script>
import axios from '../axios';

export default {
  data() {
    return {
      formData: {
        username: '',
        password: '',
        number: '',
      },
      err: '',
    };
  },
  methods: {
    async signup() {
      const { formData } = this;
      const validatePassword = /((?=.*\d)|(?=.*\W+))(?![.\n])(?=.*[A-Z])(?=.*[a-z]).*$/;
      if (formData.number === '' || formData.username === '' || formData.password === '') {
        this.err = 'Provide valid details!';
      } else if (formData.number.length !== 10) {
        this.err = 'Number must be 10 digits long!';
      } else if (formData.username.length <= 5) {
        this.err = 'Username must be greater than 5 letters!';
      } else if (!formData.password.match(validatePassword)) {
        this.err = 'Password is too weak!';
      }

      if (this.err !== '') {
        return this.err;
      }

      const res = await axios.post('/signup', { ...formData });
      if (res.status === 201) {
        return this.$router.push('/login');
      }

      this.err = 'Something went wrong!';
      return this.err;
    },
  },
};
</script>

<style>
</style>
