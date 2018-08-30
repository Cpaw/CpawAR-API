<template>
  <article v-if="isAdmin">
    <section class="notification">
      <h2>通知追加</h2>
      <form v-on:submit.prevent="addNotification">
        <div class="field">
          <div class="label">
            <label for="title">Title:</label>
          </div>
          <div class="input">
            <input class="title" type="text" name="title" v-model="notification.title">
          </div>
        </div>
        <div class="field">
          <div class="label">
            <label for="contents">Challenge text:</label>
          </div>
          <div class="contents">
            <textarea type="text" name="contents" v-model="notification.notification" cols="60" rows="5" />
          </div>
        </div>
        <div class="submit">
          <button type="submit">Submit</button>
        </div>
      </form>
      <div class="result" v-if="isSuccess">OK</div>
      <div class="error" v-if="isFailed">Invalid data</div>
    </section>
  </article>
</template>

<script>
import {HTTP} from '../Header'

export default {
  data () {
    return {
      notification: {
        title: '',
        notification: ''
      },
      isFailed: false,
      isSuccess: false,
      isAdmin: false
    }
  },
  mounted () {
    HTTP.get(`role`,
      {
        headers: {
          'Authorization': localStorage.getItem('token')
        }
      })
      .then(response => {
        if (response.data.results === 'admin') {
          this.$data.isAdmin = true
        }
      })
      .catch(e => {
      })
  },
  methods: {
    addNotification: function () {
      HTTP.post('notifications',
        this.$data.notification,
        {
          headers: {
            'Content-Type': 'application/json',
            'Authorization': localStorage.getItem('token')
          },
          withCredentials: true
        })
        .then(response => {
          this.$data.isFailed = false
          this.$data.isSuccess = true
        })
        .catch(e => {
          this.$data.isFailed = true
        })
    }
  }
}
</script>

<style scoped>
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
  font-size: 20px;
  border: solid 2px #6688cc;
  border-radius: 5px;
  outline: 0;
}
input:focus {
  border: solid 2px #6699cc;
  border-radius: 5px;
}
textarea {
  -webkit-appearance: none;
  -moz-appearance: none;
  appearance: none;
  font-family: "a-otf-ud-shin-maru-go-pr6n";
  font-size: 20px;
  border: solid 2px #6688cc;
  border-radius: 5px;
  outline: 0;
}
button {
  position: relative;
  padding: 0.25em 0.5em;
  text-decoration: none;
  color: #fff;
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
.field {
  margin: 0 auto 5vh auto;
  padding: 5px;
  font-size: 24px;
  font-weight: 700;
  width: 80vw;
}
.contents {

}
.label {
  text-align: center;
}
.result {
  margin: 2vh auto 2vh auto;
  font-size: 24px;
  color: #ff5d86;
}
.error {
  margin: 2vh auto 2vh auto;
  font-size: 24px;
  color: #ff5d86;
}
</style>
