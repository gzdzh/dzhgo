import"./store-acbd10e3.js";import{o as e,w as t,r as n,x as r}from"./lodash-es-fd5b777b.js";import{W as i}from"./@vue.runtime-core-6173334e.js";const o={resolve:null,next:null,async set(e){try{await Promise.all(e)}catch(t){}this.resolve()},async wait(){return this.next},close(){const e=document.getElementById("Loading");e&&(e.style.display="none")}};function s(e,t){let n;for(n in t)e[n]=e[n]&&"[object Object]"===e[n].toString()?s(e[n],t[n]):e[n]=t[n];return e}function a(e,t){const n=[];return e.forEach((e=>{const r=e.split(t||"/").filter(Boolean);let i=n;r.forEach(((e,t)=>{let n=i.find((t=>t.label==e));n||(n={label:e,value:e,children:r[t+1]?[]:null},i.push(n)),n.children&&(i=n.children)}))})),n}function c(e){return l(e.substring(0,e.lastIndexOf(".")))}function l(e){let t=e.lastIndexOf("/");return t=t>-1?t:e.lastIndexOf("\\"),t<0?e:e.substring(t+1)}function u(e){return e.substring(e.lastIndexOf(".")+1)}function d(e){return e.replace(/([^-])(?:-+([^-]))/g,(function(e,t,n){return t+n.toUpperCase()}))}function f(){const e=[],t="0123456789abcdef";for(let n=0;n<36;n++)e[n]=t.substr(Math.floor(16*Math.random()),1);return e[14]="4",e[19]=t.substr(3&e[19]|8,1),e[8]=e[13]=e[18]=e[23]="-",e.join("")}function h(){const{clientHeight:e,clientWidth:t}=document.documentElement,n=navigator.userAgent.toLowerCase();let r=(n.match(/firefox|chrome|safari|opera/g)||"other")[0];(n.match(/msie|trident/g)||[])[0]&&(r="msie");let i="";i="ontouchstart"in window||-1!==n.indexOf("touch")||-1!==n.indexOf("mobile")?-1!==n.indexOf("ipad")?"pad":-1!==n.indexOf("mobile")?"mobile":-1!==n.indexOf("android")?"androidPad":"pc":"pc";let o="";switch(r){case"chrome":case"safari":case"mobile":default:o="webkit";break;case"msie":o="ms";break;case"firefox":o="Moz";break;case"opera":o="O"}const s=n.indexOf("android")>0?"android":navigator.platform.toLowerCase();let a="full";a=t<768?"xs":t<992?"sm":t<1200?"md":t<1920?"xl":"full";const c=!!navigator.userAgent.match(/\(i[^;]+;( U;)? CPU.+Mac OS X/),l="pc"===i,u=!l;return{height:e,width:t,version:(n.match(/[\s\S]+(?:rv|it|ra|ie)[\/: ]([\d.]+)/)||[])[1],type:r,plat:s,tag:i,prefix:o,isMobile:u,isIOS:c,isPC:l,isMini:"xs"===a||u,screen:a}}function m(t){const r=[],i={};t.forEach((e=>i[e.id]=e)),t.forEach((e=>{const t=i[e.parentId];t?(t.children||(t.children=[])).push(e):r.push(e)}));const o=t=>{t.map((t=>{n(t.children)&&(t.children=e(t.children,"orderNum"),o(t.children))}))};return o(r),e(r,"orderNum")}function p(e){const t=[];let r=0;return function e(i,o){i.forEach((i=>{i.id||(i.id=r++),i.parentId||(i.parentId=o),t.push(i),i.children&&n(i.children)&&e(i.children,i.id)}))}(e||[],0),t}function b(e){const t={};return e.forEach((({path:e,value:n})=>{const r=e.split("/"),i=r.slice(0,r.length-1),o=l(e).replace(".ts","");let s=t;i.forEach((e=>{s[e]||(s[e]={}),s=s[e]})),s[o]=n})),t}function x(e){return!r(i(e))}function g(e){return e&&"[object Promise]"===Object.prototype.toString.call(e)}function O(e){return t(e)?`${e}px`:e}o.next=new Promise((e=>{o.resolve=e}));export{o as L,m as a,l as b,a as c,s as d,u as e,c as f,h as g,g as h,x as i,b as m,O as p,p as r,d as t,f as u};