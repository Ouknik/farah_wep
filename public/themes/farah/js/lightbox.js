document.addEventListener('DOMContentLoaded', function() {
	const elems = document.querySelectorAll('.lightbox');
	// var instances = M.Materialbox.init(elems, options);
	M.Materialbox.init(elems, {onOpenStart: onOpenStart, onCloseEnd: onCloseEnd});
});

const onOpenStart = function() {
	const element = this.el;
	if (typeof(element.dataset.large) === "undefined") return;
	element.src = element.dataset.large;
	// console.log(element.classList);
}

const onCloseEnd = function() {
	const element = this.el;
	if (typeof(element.dataset.thumbnail) === "undefined") return;
	element.src = element.dataset.thumbnail;
}


document.addEventListener('click', function (e) {
	if (e.target.classList.contains("trigger-fulltext")) {
		document.getElementById(`${e.target.dataset.idprefix}-fulltext-${e.target.dataset.id}`).classList.remove("hide");
		document.getElementById(`${e.target.dataset.idprefix}-shorttext-${e.target.dataset.id}`).classList.add("hide");
	} else if (e.target.classList.contains("trigger-shorttext")) {
		document.getElementById(`${e.target.dataset.idprefix}-fulltext-${e.target.dataset.id}`).classList.add("hide");
		document.getElementById(`${e.target.dataset.idprefix}-shorttext-${e.target.dataset.id}`).classList.remove("hide");
	}
});