<template>
  <article>
    <section class="top">
      <h1>第２回<br>Cpaw AI Competition</h1>
    </section>
    <section class="explain">
      <h2>注意事項</h2>
      <li>不正行為や妨害行為は禁止です。</li>
      <li>故意にサーバに負荷をかけないでください。</li>
      <li>回答は、1分以内に連続して提出することはできません。</li>
    </section>
  </article>
</template>

<script>
export default {
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.top {
  background: white;
  width: 70vw;
  margin: 10vh auto 0 auto;
  padding: 10px;
  border: solid 3.15px #6699cc;
  border-radius: 10px 10px;
}
h1 {
  font-family: "a-otf-ud-shin-maru-go-pr6n";
  font-size: 72px;
  letter-spacing: 10px;
  color: #6699cc;
  margin: 0;
}
h2 {
  font-family: "a-otf-ud-shin-maru-go-pr6n";
  font-size: 36px;
  color: #6699cc;
  margin: 0 0 10px 0;
}
ul {
  text-align: left;
}
li {
  text-align: left;
  padding: 10px 0 0 2em;
  font-size: 24px;
  color: #4c4c5c;
}
a {
  color: #42b983;
}
.explain {
  background: white;
  width: 60vw;
  height: 38vh;
  margin: 10vh auto 0 auto;
  border: solid 3.15px #6699cc;
  border-radius: 10px 10px;
}
</style>
