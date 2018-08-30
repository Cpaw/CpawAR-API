<template>
  <article v-if="isAdmin">
    <h2>Answerファイルのアップロード</h2>
    <section class="submit">
      <h3>問題番号：</h3>
      <input name="id" type="number" v-model="id">
      <br><br><br><br>
      <label for="csv_file">
        Select csv file
        <input @change="selectedFile" type="file" name="file" id="csv_file" style="display:none;">
      </label>
      <p>{{ filename }}</p>
      <br>
      <div class="button">
        <div class="create">
          <button @click="create" type="submit">Create</button>
          <div id="result" v-if="isSuccess">OK</div>
          <div id="error" v-if="isError">Error</div>
        </div>
        <div class="update">
          <button @click="update" type="submit">Update</button>
          <div id="result" v-if="isSuccess">OK</div>
          <div id="error" v-if="isError">Error</div>
        </div>
      </div>
    </section>
  </article>
  <article v-else>
    <h2>ログインしてください</h2>
  </article>
</template>
<script>
import {HTTP} from '../Header'

export default {
  data () {
    return {
      uploadFile: null,
      filename: '',
      isAdmin: false,
      isSuccess: false,
      isError: false,
      accuracy: 0,
      id: 0
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
        } else {
          this.$data.isAdmin = false
        }
      })
      .catch(e => {
        this.$data.isAdmin = false
      })
  },
  methods: {
    selectedFile: function (e) {
      e.preventDefault()
      let files = e.target.files
      this.$data.uploadFile = files[0]
      this.$data.filename = files[0].name
    },
    create: function () {
      this.$data.isSubmitted = true
      let formData = new FormData()
      formData.append('ansFP', this.$data.uploadFile)
      formData.append('ChallengeID', this.$data.id)
      setTimeout(() => {
        HTTP.post('answer', formData,
          {
            headers: {
              'content-type': 'multipart/form-data',
              'Authorization': localStorage.getItem('token')
            }
          })
          .then(response => {
            this.isSuccess = true
            this.isError = false
          })
          .catch(e => {
            this.$data.isError = true
            this.isSuccess = false
          })
      }, 3000)
    },
    update: function () {
      this.$data.isSubmitted = true
      let formData = new FormData()
      formData.append('ansFP', this.$data.uploadFile)
      formData.append('ChallengeID', this.$data.id)
      setTimeout(() => {
        HTTP.post('challengesfile', formData,
          {
            headers: {
              'content-type': 'multipart/form-data',
              'Authorization': localStorage.getItem('token')
            }
          })
          .then(response => {
            this.isSuccess = true
            this.isError = false
          })
          .catch(e => {
            this.$data.isError = true
            this.isSuccess = false
          })
      }, 3000)
    }

  }
}
</script>
<style scoped>
article {
  background: #fff;
  border-radius: 15px;
  border: 3px solid #6699cc;
  width: 70vw;
  height: 60vh;
  margin: 10vh auto 0 auto;
}
section {
  width: 55vw;
  margin: 0 auto;
}
h2 {
  font-family: "a-otf-ud-shin-maru-go-pr6n";
  font-size: 42px;
  height: 10vh;
}
label {
  font-family: "a-otf-ud-shin-maru-go-pr6n";
  font-size: 24px;
  padding: 0.2em 1.5em;
  border: 2px solid #6699cc;
  border-radius: 15px;
}
label:hover {
  cursor: pointer;
  background: #fafafa;
}
button {
  position: relative;
  margin-top: 2em;
  padding: 0.28em 0.5em;
  text-decoration: none;
  color: #FFF;
  background: #6699cc;
  border: solid 2px #6699cc;
  border-radius: 5px;
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
.text {
  font-size: 36px;
  font-family: "a-otf-ud-shin-maru-go-pr6n";
}
.loader,
.loader:before,
.loader:after {
  background: #6699cc;
  -webkit-animation: load1 1s infinite ease-in-out;
  animation: load1 1s infinite ease-in-out;
  width: 1em;
  height: 4em;
}
.loader {
  color: #6699cc;
  text-indent: -9999em;
  margin: 8px auto;
  position: relative;
  font-size: 11px;
  -webkit-transform: translateZ(0);
  -ms-transform: translateZ(0);
  transform: translateZ(0);
  -webkit-animation-delay: -0.16s;
  animation-delay: -0.16s;
}
.loader:before,
.loader:after {
  position: absolute;
  top: 0;
  content: '';
}
.loader:before {
  left: -1.5em;
  -webkit-animation-delay: -0.32s;
  animation-delay: -0.32s;
}
.loader:after {
  left: 1.5em;
}
@-webkit-keyframes load1 {
  0%,
  80%,
  100% {
    box-shadow: 0 0;
    height: 4em;
  }
  40% {
    box-shadow: 0 -2em;
    height: 5em;
  }
}
@keyframes load1 {
  0%,
  80%,
  100% {
    box-shadow: 0 0;
    height: 4em;
  }
  40% {
    box-shadow: 0 -2em;
    height: 5em;
  }
}
#result {
  font-size: 36px;
}
.button {
  display: -webkit-flex;
  display: -moz-flex;
  display: -ms-flex;
  display: -o-flex;
  display: flex;
  flex-direction: row;
  justify-content: space-around;
}
</style>
