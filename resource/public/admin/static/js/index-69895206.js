import{i as e}from"./index.umd.min-b551c2d1.js";import{d as o,q as t,p as r,b as s,W as a,x as i}from"./@vue.runtime-core-6173334e.js";import{r as p,i as l}from"./@vue.reactivity-54707199.js";import"./axios-bb93cf32.js";import"./element-plus-cc31d756.js";import"./lodash-es-fd5b777b.js";import"./@vueuse.core-cb7ad152.js";import"./@vueuse.shared-feb01f08.js";import"./@vue.shared-0cc4c744.js";import"./@vue.runtime-dom-532828a1.js";import"./@element-plus.icons-vue-2c9f4d7f.js";import"./@popperjs.core-b696b006.js";import"./@ctrl.tinycolor-a951068a.js";import"./dayjs-f8876307.js";import"./async-validator-ff95bfc9.js";import"./memoize-one-63ab667a.js";import"./escape-html-1935ddb3.js";import"./normalize-wheel-es-3222b0a2.js";import"./@floating-ui.dom-92665be4.js";import"./@floating-ui.core-c1b3624a.js";import"./vue-21f38777.js";const m=o({name:"cl-select",props:{modelValue:[String,Number],options:{type:[Array,Object],default:()=>[]},prop:String},emits:["update:modelValue","change"],setup(o,{emit:m}){const u=e.useCrud(),n=p(),j=t((()=>l(o.options)?o.options.value:o.options));function c(e){var t;m("update:modelValue",e),m("change",e),o.prop&&(null==(t=u.value)||t.refresh({page:1,[o.prop]:""===e?void 0:e}))}return r((()=>o.modelValue),(e=>{n.value=e}),{immediate:!0}),()=>{let e;return s(a("el-select"),{modelValue:n.value,"onUpdate:modelValue":e=>n.value=e,clearable:!0,filterable:!0,onChange:c},"function"==typeof(o=e=j.value.map((e=>s(a("el-option"),e,null))))||"[object Object]"===Object.prototype.toString.call(o)&&!i(o)?e:{default:()=>[e]});var o}}});export{m as default};