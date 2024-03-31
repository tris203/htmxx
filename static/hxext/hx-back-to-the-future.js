(function () {
  function rewriteHistory(elt) {
    var historyCache = JSON.parse(localStorage.getItem("htmx-history-cache"));
    var selector = "#" + elt.getAttribute("id");
    if (selector === "#") {
      return;
    }
    var newOuterHTML = elt.outerHTML;
    var willBeRemoved =
      elt.getAttribute("remove-me") || elt.getAttribute("data-remove-me");
    if (willBeRemoved) {
      newOuterHTML = "";
    }
    for (var i = 0; i < historyCache.length; i++) {
      var content = historyCache[i].content;
      var historyFragment = document.createDocumentFragment();
      var historyWrapper = historyFragment.appendChild(
        document.createElement("history-wrapper"),
      );
      historyWrapper.innerHTML = content;
      var targets = historyFragment.querySelectorAll(selector);
      for (var j = 0; j < targets.length; j++) {
        targets[j].outerHTML = newOuterHTML;
      }
      historyCache[i].content = historyWrapper.innerHTML;
    }
    localStorage.setItem("htmx-history-cache", JSON.stringify(historyCache));
  }
  htmx.defineExtension("back-to-the-future", {
    onEvent: function (name, evt) {
      if (name === "htmx:afterSettle") {
        // console.log(evt);
        var elt = evt.detail.elt;
        rewriteHistory(elt);
      }
    },
  });
})();
