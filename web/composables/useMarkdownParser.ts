import {
  createMarkdownParser,
  // rehypeHighlight,
  // createShikiHighlighter,
} from '@nuxtjs/mdc/runtime'
// import MaterialThemePalenight from '@shikijs/themes/material-theme-palenight'
// import TsLang from '@shikijs/langs/typescript'
// import JsLang from '@shikijs/langs/javascript'
// import YamlLang from '@shikijs/langs/yaml'

export default function useMarkdownParser() {
  let parser: Awaited<ReturnType<typeof createMarkdownParser>>

  const parse = async (markdown: string) => {
    if (!parser) {
      parser = await createMarkdownParser({
        // TODO: uncomment when shiki supports fallback languages
        // rehype: {
        //   plugins: {
        //     highlight: {
        //       instance: rehypeHighlight,
        //       options: {
        //         // Pass in your desired theme(s)
        //         theme: 'material-theme-palenight',
        //         // Create the Shiki highlighter
        //         highlighter: createShikiHighlighter({
        //           bundledThemes: {
        //             'material-theme-palenight': MaterialThemePalenight,
        //           },
        //           // Configure the bundled languages
        //           bundledLangs: {
        //             yml: YamlLang,
        //             ts: TsLang,
        //             js: JsLang,
        //             typescript: TsLang,
        //             dataviewjs: JsLang,
        //           },
        //         }),
        //       },
        //     },
        //   },
        // },
      })
    }
    return parser(markdown)
  }

  return parse
}
