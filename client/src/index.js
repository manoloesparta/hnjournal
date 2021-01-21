import './styles.css'

// color = 0x0099ff

const createArticle = (url, title) => {
  return `
  <a href="${url}" target="_blank">
    <h2>${title}</h2>
  </a>`
}

const renderArticles = async () => {
  const content = document.getElementById('content')

  for(let i = 0; i < 10; i++) {
    const res = await fetch('https://hn.manoloesparta.com/random')
    const json = await res.json()
  
    const { title, URL: url } = json
    const article = createArticle(url, title)

    content.innerHTML += article
  }
}

renderArticles()