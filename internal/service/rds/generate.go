//go:generate go run -tags generate ../../generate/tags/main.go -ListTags=yes -ListTagsInIDElem=ResourceName -ListTagsOutTagsElem=TagList -ServiceTagsSlice=yes -TagOp=AddTagsToResource -TagInIDElem=ResourceName -UntagOp=RemoveTagsFromResource -UpdateTags=yes
// ONLY generate directives and package declaration! Do not add anything else to this file.

package rds