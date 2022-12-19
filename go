waw[a_] :=
 
 WolframAlpha[
   a, {{"PopularityPod:WikipediaStatsData", 1}, "TimeSeriesData"}][[
  All, 2, 1]]
  
  acts = {"Brad Pitt", "Angelina Jolie", "Nicole Kidman", 
   "Jennifer Aniston", "Katie Holmes",
   "Tom Cruise", "Blake Lively", "Ryan Reynolds", "Bradley Cooper", 
   "Jennifer Lawrence",
   "Gwyneth Paltrow", "Robert Downey Jr", "Daniel Craig", 
   "Rachel Weisz", "Scarlett Johansson"};
   MatrixPlot[corr, ColorFunction -> "Rainbow", FrameTicks -> {
   {#, #} &[Transpose[{Range[Length[#]], #} &@acts]],
   {#, #} &[
    Transpose[{Range[Length[#]], Rotate[#, Pi/2] & /@ #} &@acts]]
   }, PlotLegends -> Automatic, Mesh -> All, 
 MeshStyle -> Directive[Opacity[.2], Dashed],
 ImageSize -> 400]
 g = WeightedAdjacencyGraph[acts, am, 
   VertexLabels -> Placed["Name", Below],
   VertexSize -> Thread[acts -> vs^1], 
   GraphLayout -> "LayeredDigraphEmbedding"];
   CommunityGraphPlot[#, FindGraphCommunities[#],
   CommunityRegionStyle -> Directive[Opacity[.2], Gray],
   CommunityBoundaryStyle -> Directive[Orange, Dashed],
   Method -> "Hierarchical", ImageSize -> 500] &@g
