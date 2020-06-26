
import * as jsYaml from 'js-yaml';
import * as _ from 'lodash';
import {MESSAGES} from 'translations/messages';

let currentLang = 'en';
let translationMap = {}

_.forEach(MESSAGES.split('---'), (rawBlock, i) => {
    try {
        let block = jsYaml.safeLoad(rawBlock) || {};
        let id = block.id || block.en;
        if (id) {
            translationMap[id] = block;
        }
    } catch (err) {
        console.info(`Failed to parse yaml block #${i+1}: ${err.toString()}\n---\n${rawBlock}`);
    }
});

export function setLang(lang: string) {
    currentLang = lang || 'en';
}

export function trEn(text: string, ...args: string[]) {
    return trLang("en", text, ...args);
}

export function trRo(text: string, ...args: string[]) {
    return trLang("ro", text, ...args);
}

export function trLang(lang: string, text: string, ...args: string[]) {
    let langMap = translationMap[text] || {};
    let rv = langMap[lang || "en"] || text;

    _.forEach(args, (a: string, i: number) => {
        rv = rv.replace(new RegExp(`%${i + 1}`, 'g'), "" + a);
    })
    return rv;
}
