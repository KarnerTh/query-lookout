import{E as b,s as E}from"./index.f49fd5c2.js";const c=[];function y(e,t){return{subscribe:d(e,t).subscribe}}function d(e,t=b){let s;const o=new Set;function r(n){if(E(e,n)&&(e=n,s)){const u=!c.length;for(const a of o)a[1](),c.push(a,e);if(u){for(let a=0;a<c.length;a+=2)c[a][0](c[a+1]);c.length=0}}}function i(n){r(n(e))}function l(n,u=b){const a=[n,u];return o.add(a),o.size===1&&(s=t(r)||b),n(e),()=>{o.delete(a),o.size===0&&s&&(s(),s=null)}}return{set:r,update:i,subscribe:l}}var g;const m=((g=globalThis.__sveltekit_is6bjv)==null?void 0:g.base)??"";var k;const w=((k=globalThis.__sveltekit_is6bjv)==null?void 0:k.assets)??m,A="1701196417454",x="sveltekit:snapshot",O="sveltekit:scroll",U="sveltekit:index",p={tap:1,hover:2,viewport:3,eager:4,off:-1};function j(e){let t=e.baseURI;if(!t){const s=e.getElementsByTagName("base");t=s.length?s[0].href:e.URL}return t}function L(){return{x:pageXOffset,y:pageYOffset}}function f(e,t){return e.getAttribute(`data-sveltekit-${t}`)}const _={...p,"":p.hover};function v(e){let t=e.assignedSlot??e.parentNode;return(t==null?void 0:t.nodeType)===11&&(t=t.host),t}function N(e,t){for(;e&&e!==t;){if(e.nodeName.toUpperCase()==="A"&&e.hasAttribute("href"))return e;e=v(e)}}function V(e,t){let s;try{s=new URL(e instanceof SVGAElement?e.href.baseVal:e.href,document.baseURI)}catch{}const o=e instanceof SVGAElement?e.target.baseVal:e.target,r=!s||!!o||S(s,t)||(e.getAttribute("rel")||"").split(/\s+/).includes("external")||e.hasAttribute("download");return{url:s,external:r,target:o}}function P(e){let t=null,s=null,o=null,r=null,i=null,l=null,n=e;for(;n&&n!==document.documentElement;)o===null&&(o=f(n,"preload-code")),r===null&&(r=f(n,"preload-data")),t===null&&(t=f(n,"keepfocus")),s===null&&(s=f(n,"noscroll")),i===null&&(i=f(n,"reload")),l===null&&(l=f(n,"replacestate")),n=v(n);return{preload_code:_[o??"off"],preload_data:_[r??"off"],keep_focus:t==="off"?!1:t===""?!0:null,noscroll:s==="off"?!1:s===""?!0:null,reload:i==="off"?!1:i===""?!0:null,replace_state:l==="off"?!1:l===""?!0:null}}function h(e){const t=d(e);let s=!0;function o(){s=!0,t.update(l=>l)}function r(l){s=!1,t.set(l)}function i(l){let n;return t.subscribe(u=>{(n===void 0||s&&u!==n)&&l(n=u)})}return{notify:o,set:r,subscribe:i}}function R(){const{set:e,subscribe:t}=d(!1);let s;async function o(){clearTimeout(s);const r=await fetch(`${w}/_app/version.json`,{headers:{pragma:"no-cache","cache-control":"no-cache"}});if(r.ok){const l=(await r.json()).version!==A;return l&&(e(!0),clearTimeout(s)),l}else throw new Error(`Version check failed: ${r.status}`)}return{subscribe:t,check:o}}function S(e,t){return e.origin!==location.origin||!e.pathname.startsWith(t)}let T;function Y(e){T=e.client}const $={url:h({}),page:h({}),navigating:d(null),updated:R()};export{U as I,p as P,O as S,x as a,V as b,P as c,L as d,m as e,N as f,j as g,Y as h,S as i,T as j,y as r,$ as s,d as w};
