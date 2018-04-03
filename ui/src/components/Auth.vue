<template>
  <article>
    <section class="signin" v-if="!isSignedIn">
      <h2>Signin</h2>
      <form v-on:submit.prevent="Signin">
        <div class="field">
          <div class="label">
            <label for="username">Name: </label>
          </div>
          <div class="input">
            <input class="name" type="text" name="username" placeholder="name" v-model="signin.username">
          </div>
        </div>
        <div class="field">
          <div class="label">
            <label for="password">Pass: </label>
          </div>
          <div class="input">
            <input class="pass" type="password" name="password" placeholder="password" v-model="signin.password">
          </div>
        </div>
        <div class="submit">
          <button type="submit">Submit</button>
        </div>
      </form>
      <div class="error" v-if="SignInError">Invalid data</div>
    </section>
    <section class="signup" v-if="!isSignedIn">
      <h2>Signup</h2>
      <form v-on:submit.prevent="Signup">
        <div class="field">
          <div class="label">
            <label for="username">Name: </label>
          </div>
          <div class="input">
            <input class="name" type="text" name="username" placeholder="name" v-model="signup.username">
          </div>
        </div>
        <div class="field">
          <div class="label">
            <label for="password">Pass: </label>
          </div>
          <div class="input">
            <input class="pass" type="password" name="password" placeholder="password" v-model="signup.password">
          </div>
        </div>
        <div class="submit">
          <button type="submit">Submit</button>
        </div>
      </form>
      <div class="error" v-if="SignUpError">Invalid data</div>
    </section>
    <section v-if="isSignedIn">
      <div class="signout">
        <a @click="Signout">Signout</a>
      </div>
    </section>
  </article>
</template>
<script>
import {HTTP} from './Header'

export default {
  data () {
    return {
      signin: {
        username: '',
        password: ''
      },
      signup: {
        username: '',
        password: ''
      },
      isSignedIn: false,
      SignInError: false,
      SignUpError: false
    }
  },
  mounted () {
    HTTP.get('role',
      {
        headers: {
          'Authorization': localStorage.getItem('token')
        }
      })
      .then(response => {
        this.$data.isSignedIn = true
      })
      .catch(e => {
        HTTP.get('signin')
          .then(response => {
            localStorage.setItem('token', response.headers.authorization)
          })
      })
  },
  methods: {
    Signin: function () {
      HTTP.post('signin',
        this.$data.signin,
        {
          headers: {
            'Content-Type': 'application/json',
            'Authorization': localStorage.getItem('token')
          },
          withCredentials: true
        })
        .then(response => {
          localStorage.setItem('token', response.headers.authorization)
          this.$router.push('/challenges')
        })
        .catch(e => {
          localStorage.setItem('token', '')
          this.$data.SignInError = true
        })
    },
    Signup: function () {
      HTTP.post('signup',
        this.$data.signup,
        {
          headers: {
            'Content-Type': 'application/json'
          },
          withCredentials: true
        })
        .then(response => {
          localStorage.setItem('token', response.headers.authorization)
          this.$router.push('/')
        })
        .catch(e => {
          localStorage.setItem('token', '')
          this.$data.SignUpError = true
        })
    },
    Signout: function () {
      HTTP.delete('signin',
        {
          headers: {
            'Authorization': localStorage.getItem('token')
          }
        })
        .then(response => {
          localStorage.setItem('token', '')
          this.$router.push('/')
        })
    }
  }
}
</script>
<style scoped>
h2 {
  font-size: 48px;
}
article {
  display: -webkit-flex;
  display: flex;
}
section {
  -webkit-justify-content: space-around;
  justify-content: space-around;
  margin: 0 auto;
}
label {
  width: 8vw;
}
input, button {
  -webkit-appearance: none;
  -moz-appearance: none;
  appearance: none;
  font-family: "a-otf-ud-shin-maru-go-pr6n";
  border: solid 2px #6699cc;
  border-radius: 5px;
  outline: 0;
}
input:focus {
  border: solid 2px #6699cc;
  border-radius: 5px;
}
button {
  position: relative;
  padding: 0.25em 0.5em;
  text-decoration: none;
  color: #FFF;
  background: #6699cc;
  border: solid 2px #6699cc;
  font-size: 20px;
  font-family: "a-otf-ud-shin-maru-go-pr6n";
  width: 10vw;
}
button:hover {
  cursor: pointer;
  background: #76a9dc;
}
button:active {
  background: #4679ac;
}
.signin {
  background: white;
  width: 36vw;
  height: 55vh;
  margin: 10vh auto 0 10vw;
  border: solid 3.15px #6699cc;
  border-radius: 10px 10px;
}
.signup {
  background: white;
  width: 36vw;
  height: 55vh;
  margin: 10vh 10vw 0 auto;
  border: solid 3.15px #6699cc;
  border-radius: 10px 10px;
}
.field {
  margin: 0 auto 5vh auto;
  padding: 5px;
  font-size: 24px;
  font-weight: 700;
  width: 24vw;
}
.label {
  text-align: left;
}
.name {
  font-size: 16px;
  width: 24vw;
}
.pass {
  font-size: 16px;
  width: 24vw;
}
.error {
  margin: 2vh auto 2vh auto;
  font-size: 24px;
  color: #ff5d86;
}
.signout {
  margin: 15vh auto 10vh auto;
  font-size: 67px;
  padding: 0.3em 1em 0.3em 1em;
  background: #FFF;
  border: 3px solid #6699cc;
  border-radius: 5px;
  cursor: pointer;
}
</style>
