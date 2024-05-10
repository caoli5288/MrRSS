<script lang="ts" setup>
import { reactive, onMounted } from 'vue'
import { GetRssContent } from '../../wailsjs/go/main/App'

type RssContent = {
  Title: string
  Link: string
  Time: string
  Summary: string
  Content: string
}

const rssContent = reactive({
  rssList: [] as RssContent[],
})

function getRssContent() {
  GetRssContent()
    .then((result: RssContent[]) => {
      rssContent.rssList = result
    })
}

onMounted(() => {
  getRssContent()
})
</script>

<template>
  <main>
    <button class="btn" @click="getRssContent">Get</button>
    <div id="rssContent" class="result">
      <div v-for="rss in rssContent.rssList" :key="rss.Title">
        <a :href="rss.Link" target="_blank">{{ rss.Title }}</a>
        <p>{{ rss.Summary }}</p>
        <p>{{ rss.Time }}</p>
      </div>
    </div>
  </main>
</template>

<style scoped>
main {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}

.result {
  line-height: 20px;
  margin: 1.5rem auto;
}

.btn {
  width: 60px;
  height: 30px;
  line-height: 30px;
  border-radius: 3px;
  border: none;
  margin: 0 0 0 20px;
  padding: 0 8px;
  cursor: pointer;
}

.btn:hover {
  background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
  color: #333333;
}
</style>
