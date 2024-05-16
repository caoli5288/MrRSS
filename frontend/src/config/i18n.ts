import { createI18n } from 'vue-i18n'

import enUS from '../i18n/en-US.json'
import enGB from '../i18n/en-GB.json'
import deDE from '../i18n/de-DE.json'
import frFR from '../i18n/fr-FR.json'
import esES from '../i18n/es-ES.json'
import zhCN from '../i18n/zh-CN.json'
import zhTW from '../i18n/zh-TW.json'
import jaJP from '../i18n/ja-JP.json'
import koKR from '../i18n/ko-KR.json'

const i18n = createI18n({
  legacy: false,
  locale: 'en-US',
  fallbackLocale: 'en-US',
  messages: {
    'en-US': enUS,
    'en-GB': enGB,
    'de-DE': deDE,
    'fr-FR': frFR,
    'es-ES': esES,
    'zh-CN': zhCN,
    'zh-TW': zhTW,
    'ja-JP': jaJP,
    'ko-KR': koKR,
  }
})

export default i18n