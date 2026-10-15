# Por qué WebTyp

[English](WHY.md) · **Español**

**Software que se cultiva, no que se acumula.**

Aplicaciones web full-stack en Go. Un solo lenguaje. Componentes tipados. Arquitectura simple.
Para humanos y para IA.

---

## Construir software es cultivar un jardín

Construir software siempre fue difícil. Pero ¿qué es construirlo? ¿Escribir código en un lenguaje,
dentro de archivos de texto? ¿Y si lo escribo más rápido, mi negocio gana porque entrego funciones
nuevas en minutos?

Sí y no.

Construir software es un proceso continuo. Se parece a cultivar un jardín: nunca termina. Si planto
cualquier semilla sin preguntarme qué es ni cómo afectará a las plantas que ya están, con el tiempo
el jardín se llena de maleza y se pudre. En cambio, si conozco cada planta y sé por qué está en ese
lugar, el jardín puede seguir creciendo, y yo puedo mirarlo y decir: *está así porque lo cultivé*.

Lo difícil nunca fue plantar. Lo difícil es **planificar, mantener y cuidar a largo plazo**.

## Lo que la IA cambió, y lo que no

La IA escribe software rápido. Pero optimiza la tarea que tiene delante; nadie más que tú responde
por el software que seguirá funcionando dentro de dos años. **El responsable siempre serás tú.**

Eso no cambió. Y entonces hay, como siempre, dos caminos:

- **Explorar:** jugar, probar herramientas, experimentar.
- **Construir para durar:** que ese código se transforme en algo rentable, que genere valor y mejore
  la vida de las personas.

Los dos son válidos. **WebTyp apuesta por el segundo.**

## Una sola forma correcta, y el compilador la exige

Si soy responsable de lo que se escribe, debo poder asegurar que es correcto. Pero cuando la IA
escribe a esa velocidad, revisarlo todo se vuelve imposible.

La respuesta conocida es escribir código legible para humanos y para IA: tipado, buena
arquitectura, responsabilidad única, modularidad, desacoplamiento, DRY, pruebas, documentación.
El problema es que son demasiadas indicaciones. El contexto del agente se consume explicándole
*cómo* quieres que construya, y nada lo obliga a seguirlo al pie de la letra.

Ahí entra WebTyp. Propone una forma **tipada y estricta** de construir:

- **No hay dos maneras de hacer lo mismo, solo una.** Las decisiones ya están tomadas en el
  framework.
- **Si no se cumplen, no compila**, y el error explica el motivo y cómo solucionarlo.
- **Se construye con piezas ya probadas**, tuyas o de la comunidad, cada una verificada de forma
  independiente.
- **Todo es un módulo con una sola responsabilidad.** WebTyp no lo sugiere: lo obliga.

> **Visión:** software que se cultiva, no que se acumula.
>
> **Misión:** hacer que la forma correcta de construir sea la única que compila.

---

## El problema es la complejidad acumulada

El desarrollo web moderno rara vez falla por falta de funciones. Falla bajo capas: un framework
de UI, un meta-framework, un bundler, un transpilador, un gestor de paquetes, una capa de tipos
sobre JavaScript, componentes de servidor y de cliente, y un archivo de configuración para cada
cosa. Cada capa tiene sentido por separado. Juntas, son un ecosistema entero que aprender antes de
la primera pantalla útil.

Frameworks como Next.js resuelven mucho, y lo hacen bien. WebTyp hace otra apuesta:

> **WebTyp no intenta ganar agregando herramientas. Intenta ganar quitándolas.**

## Cuatro promesas

### 1. Un solo lenguaje

Backend, frontend y lógica compartida se escriben en Go. Tipos, modelos y validaciones se declaran
una vez ([`model`](https://github.com/webtyp/model), [`ormc`](https://github.com/webtyp/ormc)) y se
usan en ambos lados.

Tú no escribes JavaScript. Es una promesa sobre **tu código**, no sobre el navegador: WebTyp usa el
poco JavaScript que la plataforma exige, para que tú no tengas que hacerlo.

### 2. Componentes tipados y CSS sostenible

Un componente ([`components`](https://github.com/webtyp/components)) es un tipo de Go: sus entradas
son campos, su salida es un árbol DOM ([`dom`](https://github.com/webtyp/dom)), sus eventos son
handlers tipados. Cada componente es dueño de su estructura, su estilo y su comportamiento, y el
compilador verifica cómo encajan entre sí. Sin JSX, sin hooks, sin separar componentes de cliente y
de servidor.

El CSS también es código Go tipado ([`widget/style`](https://github.com/webtyp/widget/tree/main/style)): espaciado, colores y tipografía salen de *design tokens*, no de
números mágicos ni de cientos de clases utilitarias pegadas en el HTML. Si falta una receta de
estilo, es un defecto del sistema de diseño, no una excusa para improvisar.

### 3. Arquitectura simple

Convención sobre configuración, **sin magia**. La estructura del proyecto es la configuración: no
hay configuración de build que escribir ni bundler que ajustar. Cada convención es Go explícito que
puedes abrir y seguir. → [skill `project-layout`](https://github.com/webtyp/devskills/blob/main/skills/project-layout/SKILL.md)
· [skill `app-modules`](https://github.com/webtyp/devskills/blob/main/skills/app-modules/SKILL.md)

Cada nueva función paga un *impuesto de complejidad*: ¿cuántos conceptos, archivos y ajustes nuevos
te pide? Si es más de uno de cada uno, necesita una muy buena razón para existir.

### 4. Humano o IA

La misma arquitectura sirve a un desarrollador en el teclado y a un agente trabajando por
[MCP](https://modelcontextprotocol.io). La TUI (`webtyp dev`) y el daemon MCP (`webtyp mcp`,
sobre [`mcp`](https://github.com/webtyp/mcp)) operan
el mismo proyecto con las mismas operaciones; no son dos sistemas que se distancian.

La API está pensada para que **un desarrollador junior pueda construir sin IA**. El código que una
persona puede escribir bien es también el que un agente escribe bien: las firmas tipadas lo guían y
el compilador rechaza lo incorrecto. → [skill `api-design`](https://github.com/webtyp/devskills/blob/main/skills/api-design/SKILL.md)

---

## ¿Por qué Go?

Las elecciones se hicieron con cuatro criterios:

- **Compilación rápida.** WebTyp es full-stack y construye interfaces con tecnologías web: cada
  cambio debe verse en el navegador en segundos.
- **Binarios pequeños para el navegador.** El código del frontend viaja al cliente como
  WebAssembly compilado con [TinyGo](https://tinygo.org); cada kilobyte cuenta, por eso hay
  reemplazos livianos de la librería estándar ([`fmt`](https://github.com/webtyp/fmt),
  [`json`](https://github.com/webtyp/json), [`time`](https://github.com/webtyp/time)).
- **Tipado estático.** El error debe aparecer al escribir, no al ejecutar.
- **Pocas formas de hacer lo mismo.** La sintaxis simple de Go encaja con la tesis de "una sola
  manera".

**¿Por qué no Rust?** Es excelente, pero compila lento, y ese ciclo de espera rompe el flujo de
trabajo de una interfaz web. Su expresividad también abre más caminos para resolver un mismo
problema.

**¿Por qué no Python?** Su tipado es opcional: el error llega al ejecutar, no al escribir. Y llevar
su runtime y sus librerías al navegador significa mucho binario para el cliente.

**¿Por qué no TypeScript?** Porque es justamente el problema que WebTyp quiere quitar: más
herramientas encima de JavaScript y, en la práctica, dos formas de hacer lo mismo, una en el
frontend y otra en el backend.

---

## Para quién es

**Primero que nada: freelancers, agencias y equipos pequeños (1–5 desarrolladores)** que entregan
aplicaciones web a clientes y prefieren entregar la aplicación antes que armar una cadena de
herramientas JavaScript para ella.

Lo que WebTyp busca hacer especialmente bien:

- **Progressive Web Apps:** instalables, rápidas y capaces de funcionar sin conexión.
- **Aplicaciones de negocio:** sistemas administrativos, dashboards, herramientas internas,
  formularios y CRUD sobre una base de datos real.
- **SaaS pequeños:** un solo código desde el modelo hasta el despliegue.

**Te sirve si:**

- quieres un solo lenguaje y un solo modelo mental para toda la aplicación;
- valoras pocas convenciones firmes por sobre la flexibilidad ilimitada;
- trabajas con asistentes de IA y quieres que construyan dentro de una estructura, no que inventen
  una cada vez;
- prefieres leer Go antes que configurar herramientas.

**No te sirve si:**

- necesitas React, Vue, Svelte o Angular: WebTyp no los soportará, por diseño;
- dependes del ecosistema npm o de pipelines de frontend complejos (SASS, PostCSS, module
  federation);
- necesitas hoy un ecosistema maduro, con respuestas en cada foro.

## Lo que funciona hoy

- **`webtyp dev`:** una TUI con logs en vivo, recarga en caliente y un navegador gestionado
  ([`devbrowser`](https://github.com/webtyp/devbrowser)) que se recarga con cada cambio.
- **`webtyp mcp`:** un daemon MCP al que cualquier agente puede conectarse (Claude, Cursor,
  Copilot…) para ver la interfaz, leer logs, inspeccionar el navegador, consultar la base de datos
  ([`sqlmcp`](https://github.com/webtyp/sqlmcp)) y controlar la compilación.
- **Componentes verificados en memoria:** un componente web se renderiza y se comprueba sin escribir
  un solo archivo en disco ([`sitec`](https://github.com/webtyp/sitec)).
- **Sincronización de base de datos:** modelos y esquema van al mismo paso durante el desarrollo
  ([`orm`](https://github.com/webtyp/orm), [`ddl`](https://github.com/webtyp/ddl)).
- **Despliegue en Cloudflare Workers:** el camino del proyecto a producción
  ([`goflare`](https://github.com/webtyp/goflare); ejemplo funcionando: [`goflare-demo`](https://github.com/webtyp/goflare-demo)).
- **Auto-actualización:** las nuevas versiones se instalan en segundo plano; tú decides cuándo
  reiniciar ([`update`](https://github.com/webtyp/update)).

## Antes de adoptarlo, con honestidad ❌

- **Desarrollo activo, por una sola persona.** Hoy soy el único desarrollador ([@cdvelop](https://github.com/cdvelop)). Las APIs todavía se
  están simplificando; espera cambios incompatibles y bordes ásperos.
- **Probado en un proyecto real, antes que en público.** En lugar de liberar pronto y con errores,
  construyo dos proyectos a la vez: el framework y la aplicación que lo consume, la nueva versión
  de esa plataforma de la clínica, con varios módulos. Cada error que aparece ahí corrige el
  framework antes de que llegue a ti.
- **Ecosistema joven.** Menos componentes, librerías, tutoriales y respuestas que cualquier
  framework JavaScript establecido. Puede faltarte una pieza.
- **Opinado a propósito.** WebTyp tiene una sola forma de hacer las cosas. Las contribuciones deben
  seguir esa arquitectura; eso es lo que lo mantiene simple.
- **Go en el navegador tiene límites.** WebAssembly tiene restricciones reales; algunas APIs del
  navegador críticas para el rendimiento se resuelven mejor con un poco de JavaScript, que WebTyp
  mantiene dentro del framework y no en tu código.
- **Un solo destino de despliegue por ahora:** Cloudflare Workers.

---

## De dónde viene

Hace años construí una plataforma para una clínica: backend en Go, frontend en HTML, CSS y
JavaScript, sin framework. Hoy, en 2026, sigue funcionando todos los días.

Pero su código fuente ya no existe. No usaba git, y solo queda el ejecutable. Ahora mis clientes
necesitan más: un sitio web, agentes que respondan las preguntas de siempre sobre servicios,
valores y médicos. No puedo hacer crecer un jardín del que perdí el mapa.

Esa fue la lección que dio origen a WebTyp. Que algo funcione hoy no alcanza: tiene que poder
entenderse, cuidarse y crecer mañana. Para reconstruir esa plataforma, esta vez con Go también en el
navegador, empecé a construir el framework que me hubiera gustado tener.

Tuvo varios nombres mientras encontraba su forma: **golite**, **godev**, **tinywasm** y, finalmente,
**webtyp**: *web tipada*. Cada nombre fue una versión de la misma pregunta: ¿cómo se construye
software web que dure?

## La medida del éxito

WebTyp tiene éxito cuando alguien que no es su autor puede construir una aplicación real **sin tener
que preguntar cómo funciona por dentro**, y cuando hay desarrolladores que se ganan la vida
construyendo software con él.

**¿Quieres saber cuándo esté listo?** Sigue las versiones en
**[github.com/webtyp/app](https://github.com/webtyp/app)**: *Watch → Custom → Releases*.
