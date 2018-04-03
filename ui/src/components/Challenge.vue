<template>
  <article v-if="isSignedIn">
    <h2>{{ challenge.title }}</h2>
    <section>
      <p class="text">{{ challenge.questiontext }}</p>
    </section>
    <section class="submit">
      <label for="csv_file">
        Select csv file
        <input @change="selectedFile" type="file" name="file" id="csv_file" style="display:none;">
      </label>
      <p>{{ filename }}</p>
      <br>
      <button @click="upload" type="submit" v-if="!isSubmitted">Submit</button>
      <div class="loader" v-if="isSubmitted && !(isScored || isError)">Loading...</div>
      <div id="result" v-if="isScored"><p>正解率: {{ accuracy }} %</p></div>
      <div id="error" v-if="isError"><p>ファイル形式が違うか、解答間隔が短すぎます。</p></div>
    </section>
  </article>
  <article v-else>
    <h2>ログインしてください</h2>
  </article>
</template>
<script>
import {HTTP} from './Header'

export default {
  data () {
    return {
      challenge: {
        id: 0,
        title: '',
        created_at: '',
        updated_at: '',
        questiontext: '',
        Weight: ''
      },
      uploadFile: null,
      filename: '',
      isSignedIn: false,
      isSubmitted: false,
      isScored: false,
      isError: false,
      accuracy: 0
    }
  },
  mounted () {
    HTTP.get(`challenges/` + this.$route.params.challenge_id,
      {
        headers: {
          'Authorization': localStorage.getItem('token')
        }
      })
      .then(response => {
        this.$data.isSignedIn = true
        this.$data.challenge = response.data.results.challenge
      })
      .catch(e => {
        this.$data.isSignedIn = false
      })
  },
  methods: {
    selectedFile: function (e) {
      e.preventDefault()
      let files = e.target.files
      this.$data.uploadFile = files[0]
      this.$data.filename = files[0].name
    },
    upload: function () {
      this.$data.isSubmitted = true
      let formData = new FormData()
      formData.append('ansFP', this.$data.uploadFile)
      formData.append('ChallengeID', this.$route.params.challenge_id)
      setTimeout(() => {
        HTTP.post('submit', formData,
          {
            headers: {
              'content-type': 'multipart/form-data',
              'Authorization': localStorage.getItem('token')
            }
          })
          .then(response => {
            this.$data.accuracy = response.data.results.accuracy
            this.$data.isScored = true
          })
          .catch(e => {
            this.$data.isError = true
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
  height: 13vh;
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
</style>
